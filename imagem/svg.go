package imagem

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"thtml/utils"

	"github.com/beevik/etree"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

type ConfigsSvgToPngOrJpg struct {
	Width        float64
	Height       float64
	Color        string
	OutputFormat string
	OutputPath   string
	SvgContent   string
	StrokeColor  string
}

func (c ConfigsSvgToPngOrJpg) Do() error {
	svg, err := c.PrepareSvg()
	if err != nil {
		return err
	}

	return c.RenderSvgElement(svg)
}

func (c ConfigsSvgToPngOrJpg) TratarErros() error {
	formats := []string{"jpg", "png"}
	if c.Width <= 0 || c.Width <= 0 {
		return errors.New("Tamanho de imagem invalido")
	}

	if c.Color == "" {
		return errors.New("Cor invalida")
	}

	if !utils.Contains(formats, c.OutputFormat) {
		return errors.New("Formato de saída invalido")
	}

	if c.OutputPath == "" {
		return errors.New("Lugar de saída invalido da imagem")
	}

	return nil
}

func (c ConfigsSvgToPngOrJpg) PrepareSvg() (string, error) {
	err := c.TratarErros()
	if err != nil {
		return "", err
	}

	doc := etree.NewDocument()
	if _, err := doc.ReadFrom(strings.NewReader(c.SvgContent)); err != nil {
		return "", errors.New("Erro ao analisar o JSON")
	}

	svgElement := doc.FindElement("//svg")
	if svgElement == nil {
		return "", errors.New("O svg informado não é um svg")
	}

	ChangeWidthAndHeight(c.Width, c.Height, svgElement)
	err = ChangeColor(doc, c.Color, c.StrokeColor)
	if err != nil {
		return "", err
	}

	return GetSvgString(doc, svgElement)
}

const DPI = 96.0

func GetSvgString(doc *etree.Document, svgElement *etree.Element) (string, error) {
	var b bytes.Buffer
	if _, err := doc.WriteTo(&b); err != nil {
		return "", fmt.Errorf("erro ao escrever documento SVG para buffer: %w", err)
	}
	modifiedSVGContent := b.String()
	return modifiedSVGContent, nil
}

func (c ConfigsSvgToPngOrJpg) RenderSvgElement(content string) error {
	svgReader := bytes.NewReader([]byte(content))

	ctx, err := canvas.ParseSVG(svgReader)
	if err != nil {
		return fmt.Errorf("erro ao analisar SVG: %w", err)
	}

	file, err := os.Create(c.OutputPath)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de saída: %s", err)
	}
	defer file.Close()

	switch c.OutputFormat {
	case "jpg":
		jpegWriter := renderers.JPEG(file, canvas.Resolution(DPI))
		if err := ctx.RenderTo(jpegWriter); err != nil {
			return fmt.Errorf("erro ao renderizar para JPG: %w", err)
		}
	case "png":
		pngWriter := renderers.PNG(file, canvas.Resolution(DPI))
		if err := ctx.RenderTo(pngWriter); err != nil {
			return fmt.Errorf("erro ao renderizar para PNG: %w", err)
		}
	}

	return nil
}

func ChangeWidthAndHeight(width float64, height float64, svgElement *etree.Element) {
	svgElement.CreateAttr("width", fmt.Sprintf("%v", width))
	svgElement.CreateAttr("height", fmt.Sprintf("%v", height))
}

func ChangeColor(doc *etree.Document, color string, stroke string) error {
	foundColorElement := false

	for _, el := range doc.FindElements("//*[not(name()='defs') and not(name()='style') and not(name()='metadata')]") { // Evita elementos de definição
		if el.SelectAttr("fill") != nil {
			el.CreateAttr("fill", color)
			foundColorElement = true
		}

		if el.SelectAttr("stroke") != nil {
			el.CreateAttr("stroke", stroke)
			foundColorElement = true
		}
	}

	if !foundColorElement {
		return errors.New("Aviso: Nenhum atributo 'fill' ou 'stroke' encontrado para modificar a cor. O SVG pode não ter sido alterado visualmente na cor.")
	}

	return nil
}
