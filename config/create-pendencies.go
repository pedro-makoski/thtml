package config

import (
	"thtml/executable"
	"thtml/file"
)

func Start() error {
	allJson, err := executable.GetWithJoin("./data/all.json")
	if err != nil {
		return err
	}
	isExist := file.ThisFileExists(allJson)
	if !isExist {
		err := file.CreateFile(allJson, "{}")
		if err != nil {
			return err
		}
	}

	return nil
}
