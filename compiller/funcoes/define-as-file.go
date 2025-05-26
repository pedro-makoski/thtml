package funcoes

import (
	"thtml/compiller/querycomand"
	"thtml/utils"
)

func DefineAsFile(line string, data map[string]any) error {
	args, err := querycomand.GetQuery(map[string]querycomand.ValueOptionsNumber{
		"--": {
			Required: true,
			Number:   0,
		},
		"as": {
			Required: true,
			Number:   1,
		},
	}, line, true, data, []string{"--"})
	if err != nil {
		return err
	}

	texto := args[0]
	variavel := args[1]

	data[variavel] = utils.FileFormat(texto)

	return nil
}
