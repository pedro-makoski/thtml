package imagem

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func RedimensionaImagem(caminhoEntrada string, caminhoSaida string, largura float64, altura float64) error {
	file, err := os.Open(caminhoEntrada)
	if err != nil {
		return fmt.Errorf("erro ao abrir imagem: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("erro ao decodificar imagem: %v", err)
	}

	imagemRedimensionada := resize.Resize(uint(largura), uint(altura), img, resize.Lanczos3)

	out, err := os.Create(caminhoSaida)
	if err != nil {
		return fmt.Errorf("erro ao criar imagem de saída: %v", err)
	}
	defer out.Close()

	switch strings.ToLower(filepath.Ext(caminhoSaida)) {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(out, imagemRedimensionada, nil)
	case ".png":
		err = png.Encode(out, imagemRedimensionada)
	default:
		return fmt.Errorf("formato de saída não suportado")
	}

	if err != nil {
		return fmt.Errorf("erro ao salvar imagem: %v", err)
	}

	return nil
}
