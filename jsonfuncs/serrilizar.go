package jsonfuncs

import (
	"encoding/json"
	"fmt"
	"thtml/file"
)

func Serrializar[T any](jsonObj T) (string, error) {
	bytesOfAJSON, err := json.MarshalIndent(jsonObj, "", "  ")
	if err != nil {
		return "", err
	}

	fmt.Println("Serrielizado com sucesso!!!")
	return string(bytesOfAJSON), nil
}

func SerrializarAndWriteFile[T any](jsonObj T, path string) error {
	str, err := Serrializar(jsonObj)
	if err != nil {
		return err
	}
	file.CreateFile(path, str)
	return nil
}
