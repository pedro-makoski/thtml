package funcoes

import (
	"errors"
	"thtml/compiller/querycomand"
	"thtml/funcoes"
)

func TerminalExec(line string, data map[string]any) error {
	prefixes := []string{"--", "on"}
	comands, err := querycomand.GetQuery(map[string]querycomand.ValueOptionsNumber{
		"--": {
			Number:   0,
			Required: true,
		},
		"on": {
			Number:   1,
			Required: true,
		},
	}, line, true, data, prefixes)

	if err != nil {
		return err
	}

	path := comands[1]
	allComands, err := querycomand.GetComands(comands[0], map[string]any{}, false, prefixes)
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
