package file

import (
	"fmt"
	"os"
)

func Remove(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("Erro ao remover o arquivo ou diretório: %v, erro: %w", path, err)
	}

	return nil
}
