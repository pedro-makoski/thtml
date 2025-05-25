package create

func Create(comands []string, data map[string]any) error {
	newComands, err := GetParams(comands)
	if err != nil {
		return err
	}

	err = newComands.Configs.Funcao(newComands, data)
	return err
}
