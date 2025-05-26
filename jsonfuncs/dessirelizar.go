package jsonfuncs

import (
	"encoding/json"
	"fmt"
	"thtml/file"
)

func Dessirealizar[T any](text string) (T, error) {
	var obj T
	var nulo T
	err := json.Unmarshal([]byte(text), &obj)
	if err != nil {
		return nulo, err
	}

	fmt.Println("dessirealizado com sucesso!!!")
	return obj, nil
}

func DessirealizarELerArquivo[T any](path string) (T, error) {
	text, err := file.ReadFile(path)
	var nulo T
	if err != nil {
		return nulo, err
	}
	return Dessirealizar[T](text)
}
