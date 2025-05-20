package utils

import (
	"fmt"
	"path/filepath"
)

func GetRelativePathOfTwoPaths(pathMain string, pathToGetRelative string) (string, error) {
	relativePath, err := filepath.Rel(pathMain, pathToGetRelative)
	if err != nil {
		fmt.Println("Falha ao obter copia relativa de arquivo")
		return "", err
	}

	return relativePath, nil
}
