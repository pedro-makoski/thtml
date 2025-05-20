package jsonfuncs

import (
	"fmt"
	"thtml/file"
)

func InsertInAMap(key string, value any, path string) error {
	objInPath, err := file.ReadFile(path)
	if err != nil {
		return err
	}
	objInPathAsObj, err := Dessirealizar[map[string]any](objInPath)
	if err != nil {
		return err
	}
	objInPathAsObj[key] = value
	SerrializarAndWriteFile(objInPathAsObj, path)
	fmt.Println("Inserido no JSON com sucesso!!!")
	return nil
}

func InsertWithFunc(path string, aplicateFunc func(object map[string]any) map[string]any) error {
	objInPath, err := file.ReadFile(path)
	if err != nil {
		return err
	}
	objInPathAsObj, err := Dessirealizar[map[string]any](objInPath)
	if err != nil {
		return err
	}
	objInPathAsObj = aplicateFunc(objInPathAsObj)
	err = SerrializarAndWriteFile(objInPathAsObj, path)
	if err != nil {
		return err
	}
	fmt.Println("Inserido no JSON com sucesso!!!")
	return nil
}
