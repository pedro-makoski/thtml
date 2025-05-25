package funcoes

import (
	"errors"
	"fmt"
	"thtml/compiller/querycomand"
)

func Say(comands string, data map[string]any) error {
	base, err := querycomand.GetComands(comands, data, true, []string{"--"})
	if err != nil {
		return err
	}

	if len(base) < 1 {
		return errors.New("Falta parâmetros na função Say")
	}

	text := base[0]

	fmt.Println(text)
	return nil
}
