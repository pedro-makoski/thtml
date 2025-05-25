package querycomand

import (
	"fmt"
	"strings"
	"thtml/funcoes"
	"thtml/utils"
)

func Concat(text string, data map[string]any, needToReplace bool) (string, error) {
	res := ""
	isOpenAspasDuplas := false
	isOpenAspasSimples := false
	cache := ""

	text = strings.TrimSpace(text)

	for index, actual := range text {
		if (actual == ' ' || actual == '+') && !isOpenAspasDuplas && !isOpenAspasSimples {
			if actual == ' ' && (isOpenAspasDuplas || isOpenAspasSimples) {
				cache += string(actual)
			}
			continue
		}

		if actual == '"' {
			if isOpenAspasDuplas {
				res += funcoes.ReplaceProperties(cache, data)
				cache = ""
			}
			isOpenAspasDuplas = !isOpenAspasDuplas
			continue
		}

		if actual == '\'' {
			if isOpenAspasSimples {
				res += funcoes.ReplaceProperties(cache, data)
				cache = ""
			}

			isOpenAspasSimples = !isOpenAspasSimples
			continue
		}

		if !isOpenAspasDuplas && !isOpenAspasSimples && (index == len(text)-1 || text[index+1] == '+') {
			variavel := cache + string(actual)
			value := variavel
			var ok bool
			if needToReplace {
				value, ok = data[variavel].(string)
				if !ok {
					return "", fmt.Errorf("variável %s não encontrada, tente colocar %s entre aspas assim: '%s'", variavel, variavel, variavel)
				}

			}

			res += value
			continue
		}

		cache += string(actual)
	}

	return res, nil
}

func ConcatAll(text []string, data map[string]any, needToReplace bool, posesToReplace []string) ([]string, error) {
	concat := text

	for index, actual := range text {
		isPossible := false
		if index == 0 {
			isPossible = utils.Contains(posesToReplace, "--")
		} else {
			isPossible = utils.Contains(posesToReplace, text[index-1])
		}
		concated, err := Concat(actual, data, needToReplace && (len(posesToReplace) == 0 || isPossible))
		if err != nil {
			return nil, err
		}

		concat[index] = concated
	}

	return concat, nil
}

func GetComandWithoutConcat(comandsInString string, data map[string]any) ([]string, error) {
	comands := []string{}

	contInList := 0
	isOpenAspasDuplas := false
	isOpenAspasSimples := false
	addedConcat := false
	writedInConcat := false
	wasAdded := false

	comands = append(comands, "")
	for index, actual := range comandsInString {
		before := ""
		if index != 0 {
			before = string(comandsInString[index-1])
		}
		if actual != '+' && !isOpenAspasDuplas && !isOpenAspasSimples && actual != ' ' && !addedConcat && index != 0 && len(comands[contInList]) != 0 && (before == "+" || before == " ") {
			contInList++
			comands = append(comands, "")
			comands[contInList] += string(actual)
			wasAdded = true
		}

		if actual == '"' {
			isOpenAspasDuplas = !isOpenAspasDuplas
		}

		if actual == '\'' {
			isOpenAspasSimples = !isOpenAspasSimples

		}

		if (actual == '+' && !isOpenAspasDuplas && !isOpenAspasSimples) || (addedConcat && writedInConcat && actual == ' ' && !isOpenAspasDuplas && !isOpenAspasSimples) {
			if addedConcat {
				addedConcat = false
				continue
			}

			addedConcat = true
			comands[contInList] += string(actual)
			wasAdded = false
			continue
		}

		if !wasAdded {
			comands[contInList] += string(actual)
		}
		if actual != ' ' && addedConcat {
			writedInConcat = true
		}
		wasAdded = false
	}

	return comands, nil
}

func GetComands(comandsInString string, data map[string]any, needToReplace bool, posesToReplace []string) ([]string, error) {
	comands, err := GetComandWithoutConcat(comandsInString, data)
	if err != nil {
		return []string{}, nil
	}

	concated, err := ConcatAll(comands, data, needToReplace, posesToReplace)
	return concated, err
}

func GetSeparedComands(comandInString string) []string {
	return utils.SplitBy(comandInString, []string{`\s`})
}
