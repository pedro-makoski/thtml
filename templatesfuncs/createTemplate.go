package templatesfuncs

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"thtml/file"
)

func CreateTemplateFile(data map[string]any, template *template.Template, path string) error {
	err := file.CreateFolder(filepath.Dir(path))
	if err != nil {
		return err
	}

	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Erro ao criar o arquivo: %v", path)
	}
	defer outputFile.Close()

	err = template.Execute(outputFile, data)
	return err
}
