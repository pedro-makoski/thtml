package querycomand

import (
	"errors"
	"fmt"
	"thtml/utils"
)

type ValueOptionsNumber struct {
	Required bool
	Number   int
}

func GetQuery(query map[string]ValueOptionsNumber, line string, needToReplace bool, data map[string]any, posesToReplace []string) ([]string, error) {
	comands, err := GetComands(line, map[string]any{}, needToReplace, posesToReplace)
	if err != nil {
		return nil, err
	}
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
				return res, errors.New("Está faltando argumento incial")
			}
			continue
		}

		indexInComand := utils.GetIndexOfValue(comands, key)
		exists := !(indexInComand == -1 || indexInComand+1 > len(comands)-1)
		if !exists && value.Required {
			return res, fmt.Errorf("Está faltando o argumento após o %v\n", key)
		}

		if exists {
			newValue := comands[indexInComand+1]
			res[value.Number] = newValue
			continue
		}

		res[value.Number] = ""
	}

	return ConcatAll(res, data, needToReplace, posesToReplace)
}
