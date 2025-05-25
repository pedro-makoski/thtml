package create

import (
	"errors"
	"fmt"
	"regexp"
	"thtml/utils"
)

func GetParams(comands []string) (CreateParams, error) {
	res := CreateParams{}
	res.Optionals = make(map[string]any)
	res.Params = make(map[string]any)
	if len(comands) < 1 {
		return res, errors.New("Está faltando argumentos")
	}
	res.Name = comands[0]
	comandCreate, ok := GetCreateFunctions()[res.Name]
	if !ok {
		return res, fmt.Errorf("O comando create %v não existe", comandCreate)
	}

	beforeIfParam := ""
	beforeItemOptional := ""
	for index, comand := range comands[1:] {
		if index == 0 && comandCreate.HasStart {
			res.Params["--"] = comand
			continue
		}

		if beforeIfParam != "" {
			res.Params[beforeIfParam] = comand
			beforeIfParam = ""
			continue
		}

		re := regexp.MustCompile("--(.+)=(.+)")
		sub := re.FindStringSubmatch(comand)
		if len(sub) == 3 {
			prefix := sub[1]
			sufix := sub[2]
			if exists := utils.GetIndexOfValue(comandCreate.Optionals, prefix); exists != -1 {
				res.Optionals[prefix] = sufix
				continue
			}

			return res, fmt.Errorf("O comando %v não existe na função create %v", prefix, res.Name)
		}

		primordialRe := regexp.MustCompile("--(.+)")
		sub = primordialRe.FindStringSubmatch(comand)
		if len(sub) == 2 {
			prefix := sub[0]
			if exists := utils.GetIndexOfValue(comandCreate.Optionals, prefix); exists != -1 {
				beforeItemOptional = prefix
			}

			return res, fmt.Errorf("O comando %v não existe para a função create %v", prefix, res.Name)
		}

		if beforeItemOptional != "" {
			res.Optionals[beforeItemOptional] = comand
			beforeIfParam = ""
		}

		if prefix := utils.GetIndexOfValue(comandCreate.Params, comand); prefix != -1 {
			beforeIfParam = comand
			continue
		}

		return res, fmt.Errorf("O parâmetro %v não existe no create %v", comand, res.Name)
	}

	return res, nil
}
