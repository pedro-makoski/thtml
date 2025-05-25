package funcoes

import "thtml/compiller/create"

func Create(line string, data map[string]any) error {
	return create.Create(line, data)
}
