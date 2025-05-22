package create

import "fmt"

func Create(comands []string) error {
	newComands, err := GetParams(comands)
	if err != nil {
		return err
	}

	fmt.Println(newComands)
	return nil
}
