package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func RenameFile(path string, newBaseName string) error {
	dir := filepath.Dir(path)
	err := os.Rename(path, filepath.Join(dir, newBaseName))
	if err != nil {
		fmt.Println("Falha ao renomear arquivo")
		return err
	}

	return nil
}

func CopyFile(origem string, destino string) error {
	arquivoOrigem, err := os.Open(origem)
	if err != nil {
		fmt.Printf("Problemas ao abrir o arquivo %v\n", origem)
		return err
	}
	defer arquivoOrigem.Close()

	folderDestino := filepath.Dir(destino)
	CreateFolder(folderDestino)
	arquivoDestino, err := os.Create(destino)
	if err != nil {
		fmt.Printf("Problemas ao criar o arquivo %v\n", destino)
		return err
	}
	defer arquivoDestino.Close()

	_, err = io.Copy(arquivoDestino, arquivoOrigem)
	if err != nil {
		fmt.Printf("Problemas ao copiar o arquivo %v\n", origem)
		return err
	}

	return nil
}
