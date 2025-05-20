package refresh

import (
	"errors"
	"fmt"
	"thtml/executable"
	"thtml/jsonfuncs"
	"thtml/stepsfuncs"
	"thtml/templatesfuncs"
	"thtml/utils"
)

func DoRefreshEspecific(funcao func(path string) (string, error), path string) (string, error) {
	newPath, err := executable.GetWithJoin(path)
	if err != nil {
		return "", err
	}
	value, err := funcao(newPath)
	if err != nil {
		return "", err
	}
	return value, nil
}

func DoRefreshArea(obj map[string]any) (map[string]any, error) {
	newObj := utils.CopyMap(obj)

	templates, ok := newObj["templates"].(map[string]any)
	if !ok {
		err := errors.New("Existe uma má formatação no código, certifique-se que nada foi alterado erroneamente")
		return nil, err
	}

	keys, values := utils.GetKeysAndValues(templates)
	for i, key := range keys {
		value := values[i]
		valueNew := value.(string)
		delete(templates, key)
		newKey, err := DoRefreshEspecific(templatesfuncs.GetDefineNameId, valueNew)
		if err != nil {
			return nil, err
		}
		templates[newKey] = value
	}

	steps, ok := newObj["steps"].(map[string]any)
	if !ok {
		return nil, errors.New("Existe uma má formatação no código, certifique-se que nada foi alterado erroneamente")
	}

	keys, values = utils.GetKeysAndValues(steps)
	for i, key := range keys {
		value := values[i]
		valueNew := value.(string)
		delete(steps, key)
		newKey, err := DoRefreshEspecific(stepsfuncs.GetComandName, valueNew)
		if err != nil {
			return nil, err
		}
		steps[newKey] = value
	}

	return newObj, nil
}

func RefreshAll() error {
	path, err := executable.GetWithJoin("./data/all.json")
	if err != nil {
		return err
	}
	obj, err := jsonfuncs.DessirealizarELerArquivo(path)
	if err != nil {
		return err
	}
	for key, value := range obj {
		valueConverted, ok := value.(map[string]any)
		if !ok {
			err := errors.New("Existe uma má formatação no JSON, certifique-se que nada foi alteado erroneamente")
			return err
		}
		value, err = DoRefreshArea(valueConverted)
		if err != nil {
			return err
		}
		obj[key] = value
	}
	err = jsonfuncs.SerrializarAndWriteFile(obj, path)
	if err != nil {
		return err
	}
	fmt.Println("Atualizado com sucesso!!!")
	return nil
}
