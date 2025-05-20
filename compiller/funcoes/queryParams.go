package funcoes

import (
	"errors"
	"thtml/utils"
)

type ValueOptionsNumber struct {
	Required bool
	Number   int
}

func GetQuery(query map[string]ValueOptionsNumber, line string) ([]string, error) {
	comands := GetSeparedComands(line)
	keys, values := utils.GetKeysAndValues(query)
	res := make([]string, len(keys))
	for index, key := range keys {
		value := values[index]
		if key == "--" {
			if len(comands) >= 1 {
				res[value.Number] = comands[0]
				continue
			}

			if value.Required {
				return res, errors.New("Está faltando argumentos")
			}
			continue
		}

		indexInComand := utils.GetIndexOfValue(comands, key)
		exists := !(indexInComand == -1 || indexInComand+1 > len(comands)-1)
		if !exists && value.Required {
			return res, errors.New("Está faltando argumentos")
		}

		if exists {
			newValue := comands[indexInComand+1]
			res[value.Number] = newValue
			continue
		}

		res[value.Number] = ""
	}

	return res, nil
}
