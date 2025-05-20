package funcoes

import (
	"fmt"
)

func Say(comands string, data map[string]any) error {
	base, err := getComands(comands, data)
	if err != nil {
		return err
	}

	text := base[0]

	fmt.Println(text)
	return nil
}
