package funcoes

import "thtml/compiller/create"

func Create(line string, data map[string]any) error {
	comands, err := getComands(line, data, false)
	if err != nil {
		return err
	}

	return create.Create(comands)
}
