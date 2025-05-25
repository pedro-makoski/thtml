package funcoes

import (
	"errors"
	"thtml/compiller/querycomand"
)

func Define(line string, data map[string]any) error {
	prefixes := []string{"--", "as"}
	comands, err := querycomand.GetQuery(map[string]querycomand.ValueOptionsNumber{
		"--": {
			Required: true,
			Number:   0,
		},
		"as": {
			Required: true,
			Number:   1,
		},
	}, line, true, data, prefixes)
	if err != nil {
		return err
	}
	if len(comands) < 2 {
		return errors.New("Está faltando argumentos para definir")
	}

	variableName := comands[0]
	text := comands[1]

	comand, err := querycomand.GetComands(text, data, false, prefixes)
	if err != nil {
		return err
	}

	data[variableName] = comand[0]

	return nil
}
