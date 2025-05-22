package funcoes

import (
	"errors"
	"thtml/funcoes"
)

func TerminalExec(line string, data map[string]any) error {
	comands, err := GetQuery(map[string]ValueOptionsNumber{
		"--": {
			Number:   0,
			Required: true,
		},
		"on": {
			Number:   1,
			Required: true,
		},
	}, line, false, data)

	if err != nil {
		return err
	}

	path := comands[1]
	allComands, err := getComands(comands[0], map[string]any{}, false)
	if err != nil {
		return err
	}

	if len(allComands) < 1 {
		return errors.New("Está faltando o comando")
	}
	name := allComands[0]
	args := []string{""}
	if len(allComands) > 1 {
		args = allComands[1:]
	}

	err = funcoes.TerminalExec(path, name, args...)

	return err
}
