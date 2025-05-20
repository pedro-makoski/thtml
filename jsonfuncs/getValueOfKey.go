package jsonfuncs

import (
	"errors"
	"thtml/utils"
)

func GetValueOfKey[T any](path string, keys ...string) (T, error) {
	json, err := DessirealizarELerArquivo(path)
	var nulo T
	if err != nil {
		return nulo, err
	}

	value, err := utils.GetValueOfKeys(json, keys...)
	if err != nil {
		return nulo, err
	}

	newValue, ok := value.(T)
	if !ok {
		return nulo, errors.New("Houve uma má formatação do arquivo")
	}
	return newValue, nil
}
