package funcoes

import (
	"errors"
)

func Define(line string, data map[string]any) error {
	comands, err := GetQuery(map[string]ValueOptionsNumber{
		"--": {
			Required: true,
			Number:   0,
		},
		"as": {
			Required: true,
			Number:   1,
		},
	}, line, false)
	if err != nil {
		return err
	}
	if len(comands) < 2 {
		return errors.New("Está faltando argumentos para definir")
	}

	variableName := comands[0]
	text := comands[1]

	comand, err := getComands(text, data, false)
	if err != nil {
		return err
	}

	data[variableName] = comand[0]

	return nil
}
