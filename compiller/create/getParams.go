package create

import (
	"errors"
	"fmt"
	"regexp"
	"thtml/compiller/create/estruturas"
	"thtml/compiller/querycomand"
	"thtml/utils"
)

func GetParams(comandsString string, data map[string]any) (estruturas.CreateParams, error) {
	comands, err := querycomand.GetComandWithoutConcat(comandsString, data)
	if err != nil {
		return estruturas.CreateParams{}, nil
	}
	res := estruturas.CreateParams{}
	res.Optionals = make(map[string]any)
	res.Params = make(map[string]any)
	if len(comands) < 1 {
		return res, errors.New("Está faltando argumentos")
	}
	res.Name = comands[0]
	comandCreate, ok := GetCreateFunctions()[res.Name]
	if !ok {
		return res, fmt.Errorf("O comando create %v não existe", res.Name)
	}

	toSubstituteValue := comandCreate.ValoresSubstituiveis
	comands, err = querycomand.ConcatAll(comands[1:], data, true, toSubstituteValue)
	if err != nil {
		return estruturas.CreateParams{}, err
	}

	beforeIfParam := ""
	beforeItemOptional := ""
	for index, comand := range comands {
		if index == 0 && comandCreate.HasStart {
			res.Start = comand
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

	res.Configs = comandCreate
	return VerifyObrigatoriosAndDefaults(res)
}

func VerifyObrigatoriosAndDefaults(comandos estruturas.CreateParams) (estruturas.CreateParams, error) {
	newComandos := comandos
	chavesInfomadas, _ := utils.GetKeysAndValues(comandos.Params)
	for _, obrigatorio := range comandos.Configs.ParametrosObrigatorios {
		param := comandos.Configs.Params[obrigatorio]
		if exists := utils.GetIndexOfValue(chavesInfomadas, param); exists == -1 {
			return estruturas.CreateParams{}, fmt.Errorf("Está faltando o parâmetro: %v", param)
		}
	}

	defaultsPoses, defaultsValue := utils.GetKeysAndValues(comandos.Configs.Defaults)

	for idx, posOfDefault := range defaultsPoses {
		defaultValue := defaultsValue[idx]

		param := comandos.Configs.Params[posOfDefault]
		if exists := utils.GetIndexOfValue(chavesInfomadas, param); exists == -1 {
			newComandos.Params[param] = defaultValue
		}
	}

	return newComandos, nil
}
