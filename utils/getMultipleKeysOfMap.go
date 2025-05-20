package utils

import (
	"errors"
	"fmt"
)

func GetValueOfKeys[K comparable](obj map[K]any, keys ...K) (any, error) {
	var value any
	var ok bool

	for index, key := range keys {
		value, ok = obj[key]
		if !ok {
			return value, fmt.Errorf("a key %v não foi encontrada", key)
		}

		if index == len(keys)-1 {
			break
		}
		obj, ok = value.(map[K]any)
		if !ok {
			return value, errors.New("Houve uma má formatação do arquivo")
		}
	}

	return value, nil
}
