package file

import (
	"os"
	"path/filepath"
)

func CopyDir(origem string, destino string) error {
	err := os.MkdirAll(destino, os.ModePerm)
	if err != nil {
		return err
	}

	origemDir, err := os.Open(origem)
	if err != nil {
		return err
	}
	defer origemDir.Close()

	entradas, err := origemDir.Readdir(-1)
	if err != nil {
		return err
	}

	for _, entrada := range entradas {
		srcPath := filepath.Join(origem, entrada.Name())
		dstPath := filepath.Join(destino, entrada.Name())

		if entrada.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
			continue
		}

		err = CopyFile(srcPath, dstPath)
		if err != nil {
			return err
		}
	}

	return nil
}
