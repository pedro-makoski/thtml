package imagem

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"thtml/file"
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

	if !utils.Contains(formats, c.OutputFormat) {
		return errors.New("Formato de saída invalido")
	}

	if c.OutputPath == "" {
		return errors.New("Lugar de saída invalido da imagem")
	}

	return nil
}

func (c *ConfigsSvgToPngOrJpg) PrepareSvg() (string, error) {
	err := c.TratarErros()
	if err != nil {
		return "", err
	}

	doc := etree.NewDocument()
	if _, err := doc.ReadFrom(strings.NewReader(c.SvgContent)); err != nil {
		return "", fmt.Errorf("Erro ao analisar o svg: %v", c.SvgContent)
	}

	svgElement := doc.FindElement("//svg")
	if svgElement == nil {
		return "", fmt.Errorf("O svg informado %v não é um svg", c.SvgContent)
	}

	NormalizeViewBox(svgElement)

	ChangeWidthAndHeight(c.Width, c.Height, svgElement)
	err = ChangeColor(doc, c.Color, c.StrokeColor)
	if err != nil {
		return "", err
	}

	c.OutputPath = file.ChangeExt(c.OutputPath, c.OutputFormat)

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
	canv, err := canvas.ParseSVG(bytes.NewReader([]byte(content)))
	if err != nil {
		return fmt.Errorf("erro ao analisar SVG: %w", err)
	}

	dpiOption := canvas.DPMM(DPI / 25.4)

	if err := renderers.Write(c.OutputPath, canv, dpiOption); err != nil {
		return fmt.Errorf("erro ao salvar %s: %w", c.OutputFormat, err)
	}

	RedimensionaImagem(c.OutputPath, c.OutputPath, c.Width, c.Height)
	return nil
}

func NormalizeViewBox(svg *etree.Element) {
	vb := svg.SelectAttrValue("viewBox", "")
	parts := strings.Fields(vb)
	if len(parts) != 4 {
		return
	}
	minX, _ := strconv.ParseFloat(parts[0], 64)
	minY, _ := strconv.ParseFloat(parts[1], 64)
	origW, _ := strconv.ParseFloat(parts[2], 64)
	origH, _ := strconv.ParseFloat(parts[3], 64)

	svg.RemoveAttr("viewBox")
	svg.CreateAttr("viewBox", fmt.Sprintf("0 0 %v %v", origW, origH))

	g := etree.NewElement("g")
	g.CreateAttr("transform", fmt.Sprintf("translate(%v %v)", -minX, -minY))

	for _, child := range svg.ChildElements() {
		svg.RemoveChild(child)
		g.AddChild(child)
	}
	svg.AddChild(g)
}

func ChangeWidthAndHeight(width, height float64, svg *etree.Element) {
	vb := svg.SelectAttrValue("viewBox", "")
	parts := strings.Fields(vb)
	if len(parts) == 4 {
		origW, _ := strconv.ParseFloat(parts[2], 64)
		origH, _ := strconv.ParseFloat(parts[3], 64)
		aspect := origW / origH

		switch {
		case width > 0 && height > 0:
			height = width / aspect
		case width > 0:
			height = width / aspect
		case height > 0:
			width = height * aspect
		}
	}

	if width > 0 {
		svg.RemoveAttr("width")
		svg.CreateAttr("width", fmt.Sprintf("%.0f", width))
	}
	if height > 0 {
		svg.RemoveAttr("height")
		svg.CreateAttr("height", fmt.Sprintf("%.0f", height))
	}

	svg.RemoveAttr("preserveAspectRatio")
	svg.CreateAttr("preserveAspectRatio", "xMidYMid meet")
}

func removeStyleProperty(styleString, propertyToRemove string) (newStyleString string, removed bool) {
	properties := strings.Split(styleString, ";")
	var newProperties []string
	propertyFound := false

	for _, prop := range properties {
		prop = strings.TrimSpace(prop)
		if prop == "" {
			continue
		}
		parts := strings.SplitN(prop, ":", 2)
		if len(parts) == 2 && strings.TrimSpace(parts[0]) == propertyToRemove {
			propertyFound = true
		} else {
			newProperties = append(newProperties, prop)
		}
	}

	if !propertyFound {
		return styleString, false
	}

	return strings.Join(newProperties, "; "), true
}

func ChangeColor(doc *etree.Document, fillColor, strokeColor string) error {
	root := doc.Root()
	if root == nil {
		return errors.New("documento XML vazio")
	}

	if fillColor != "" {
		fillColor = strings.ToLower(fillColor)
	}
	if strokeColor != "" {
		strokeColor = strings.ToLower(strokeColor)
	}

	for _, el := range root.FindElements(".//*") {
		switch el.Tag {
		case "defs", "style", "metadata", "script", "title", "desc":
			continue
		}

		if fillColor != "" {
			el.RemoveAttr("fill")
			el.CreateAttr("fill", fillColor)

			if styleAttr := el.SelectAttr("style"); styleAttr != nil {
				newStyle, removed := removeStyleProperty(styleAttr.Value, "fill")
				if removed {
					if strings.TrimSpace(newStyle) == "" {
						el.RemoveAttr("style")
					} else {
						styleAttr.Value = newStyle
					}
				}
			}
		}
		if strokeColor != "" {
			el.RemoveAttr("stroke")
			el.CreateAttr("stroke", strokeColor)

			if styleAttr := el.SelectAttr("style"); styleAttr != nil {
				newStyle, removed := removeStyleProperty(styleAttr.Value, "stroke")
				if removed {
					if strings.TrimSpace(newStyle) == "" {
						el.RemoveAttr("style")
					} else {
						styleAttr.Value = newStyle
					}
				}
			}
		}
	}

	return nil
}
