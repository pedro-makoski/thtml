package file

import (
	"fmt"
	"os"
	"path"
)

func CreateFolder(folder string) error {
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		return err
	}

	fmt.Printf("Pasta %v criada com sucesso\n", folder)
	return nil
}

func CreateFile(file string, content string) error {
	fileName := path.Base(file)
	folder := path.Dir(file)

	CreateFolder(folder)
	err := os.WriteFile(file, []byte(content), 0755)
	if err != nil {
		return err
	}

	fmt.Printf("Arquivo %v criado/editado com sucesso\n", fileName)
	return nil
}
