package controllers

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"thtml/compiller/create/estruturas"
	"thtml/file"
	"thtml/imagem"
)

func CreateImageFileFromSvg(params estruturas.CreateParams, data map[string]any) error {
	width, err := strconv.ParseFloat(params.Params["width"].(string), 64)
	if err != nil {
		return errors.New("Erro ao converter a largura no tipo numerico")
	}
	height, err := strconv.ParseFloat(params.Params["height"].(string), 64)
	if err != nil {
		return errors.New("Erro ao converter a largura no tipo numerico")
	}

	definicoes := imagem.ConfigsSvgToPngOrJpg{
		Width:        width,
		Height:       height,
		Color:        params.Params["color"].(string),
		OutputFormat: params.Params["format"].(string),
		OutputPath:   params.Params["on"].(string),
		SvgContent:   params.Start,
		StrokeColor:  params.Params["stroke"].(string),
	}

	err = file.CreateFolder(filepath.Dir(definicoes.OutputPath))
	if err != nil {
		return err
	}

	err = definicoes.Do()
	if err != nil {
		return err
	}

	fmt.Println("SVG criado com sucesso!")
	return nil
}
