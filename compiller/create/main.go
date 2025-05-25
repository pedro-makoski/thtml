package create

func Create(comands string, data map[string]any) error {
	newComands, err := GetParams(comands, data)
	if err != nil {
		return err
	}

	if newComands.Configs.Funcao != nil {
		err = newComands.Configs.Funcao(newComands, data)
	}
	return err
}
