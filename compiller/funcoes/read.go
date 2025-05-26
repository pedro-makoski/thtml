package funcoes

import (
	"fmt"
	"thtml/compiller/querycomand"
	"thtml/file"
)

func Read(comand string, data map[string]any) error {
	args, err := querycomand.GetQuery(map[string]querycomand.ValueOptionsNumber{
		"--": {
			Required: true,
			Number:   0,
		},
		"as": {
			Required: true,
			Number:   1,
		},
	}, comand, true, data, []string{"--", "as"})

	if err != nil {
		return err
	}

	filePath := args[0]
	as := args[1]

	content, err := file.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo %s: %w", filePath, err)
	}

	data[as] = content
	return nil
}
