package configonjsonpart

import (
	"errors"
	"fmt"
	"path/filepath"
	"thtml/executable"
	"thtml/file"
	"thtml/jsonfuncs"
	"thtml/local"
	"thtml/utils"
)

func StartSomething(comands []string, pathToFormatNameAndNameWithOutExt string, nameKey string, funcDefineName func(path string) (string, error)) error {
	if len(comands) < 2 {
		return errors.New("está faltando parâmetros")
	}

	pathJson, err := executable.GetWithJoin("./data/all.json")
	if err != nil {
		return err
	}

	templateFile := comands[0]
	projectName := comands[1]
	fullTemplateFile, err := local.GetJoinPathInAbsPathActual(templateFile)
	if err != nil {
		return err
	}
	templateName, err := funcDefineName(fullTemplateFile)
	if err != nil {
		return err
	}

	pathOfData, err := jsonfuncs.GetValueOfKey[string](pathJson, projectName, "path-of-data")
	if err != nil {
		return err
	}
	pathNewToTemplate := filepath.Join(pathOfData, fmt.Sprintf(pathToFormatNameAndNameWithOutExt, templateName))
	err = file.CopyFile(fullTemplateFile, pathNewToTemplate)
	if err != nil {
		return err
	}
	folder, err := executable.GetFolder()
	if err != nil {
		return err
	}
	newPathToNewTemplate, err := utils.GetRelativePathOfTwoPaths(folder, pathNewToTemplate)
	if err != nil {
		return err
	}

	if !filepath.IsAbs(newPathToNewTemplate) {
		newPathToNewTemplate, err = filepath.Abs(newPathToNewTemplate)
		if err != nil {
			return err
		}
	}

	jsonfuncs.InsertWithFunc(pathJson, func(object map[string]any) map[string]any {
		return utils.PutValueInMultipleKeys(newPathToNewTemplate, object, projectName, nameKey, templateName)
	})
	return nil
}

func GetPathOfJSON(keyName string, projectName string, whereIsAllPaths string) (string, error) {
	pathAct, err := executable.GetWithJoin("./data/all.json")
	if err != nil {
		return "", err
	}
	obj, err := jsonfuncs.DessirealizarELerArquivo[map[string]any](pathAct)
	if err != nil {
		return "", err
	}
	objOfProject, ok := obj[projectName].(map[string]any)
	if !ok {
		err := errors.New("Aconteceu uma má formatação no JSON")
		return "", err
	}

	all, ok := objOfProject[whereIsAllPaths].(map[string]any)
	if !ok {
		err := errors.New("Aconteceu uma má formatação no JSON")
		return "", err
	}

	path, ok := all[keyName].(string)
	if !ok {
		byteKey := []byte(whereIsAllPaths)
		keyNameOfAllPathsWithoutS := string(byteKey[:len(byteKey)-1])
		err := fmt.Errorf("Este %v não existe. Caso tenha certeza que exista execute: thtml refresh e tente novamente", keyNameOfAllPathsWithoutS)
		return "", err
	}

	return path, nil
}
