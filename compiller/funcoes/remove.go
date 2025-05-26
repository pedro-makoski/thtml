package funcoes

import (
	"thtml/compiller/querycomand"
	"thtml/file"
)

func Remove(line string, data map[string]any) error {
	args, err := querycomand.GetQuery(map[string]querycomand.ValueOptionsNumber{
		"--": {
			Required: true,
			Number:   0,
		},
	}, line, true, data, []string{"--", "from", "as", "on"})

	if err != nil {
		return err
	}

	path := args[0]
	err = file.Remove(path)

	return err
}
