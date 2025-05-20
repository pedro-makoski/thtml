package funcoes

import (
	"thtml/funcoes"
)

func Ask(comands string, data map[string]any) error {
	allComands, err := getComands(comands, data)

	if err != nil {
		return err
	}

	for _, toAsk := range allComands {
		if len(toAsk) == 0 {
			continue
		}
		text, err := funcoes.Read(funcoes.ReplaceProperties(toAsk, data) + ": ")
		if err != nil {
			return err
		}
		data[toAsk] = text
	}

	return nil
}
