package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func Rename(oldPath string, newPath string) error {
	destDir := filepath.Dir(newPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	if err := os.Rename(oldPath, newPath); err != nil {
		return err
	}

	fmt.Println("Arquivo renomeado com sucesso!")
	return nil

}
