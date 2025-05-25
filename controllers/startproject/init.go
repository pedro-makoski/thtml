package startproject

import (
	"errors"
	"fmt"
	"path/filepath"
	"thtml/executable"
	"thtml/file"
	"thtml/jsonfuncs"
	"thtml/local"
)

func InitProject(params []string) error {
	if len(params) < 3 {
		return errors.New("Coloque o nome do projeto e o local do projeto(-- para ficar no local do exe)")
	}

	nome := params[0]
	pathOfProject := params[1]
	pathOfData := params[2]

	localRelative := "./"
	if len(params) >= 2 {
		localRelative = pathOfProject
		if localRelative == "." {
			localRelative = "./"
		}
	}
	localAbs, err := local.GetJoinPathInAbsPathActual(localRelative)
	if err != nil {
		return err
	}

	if pathOfData == "--" {
		pathOfData, err = executable.GetWithJoin("./data/steps")
		if err != nil {
			return err
		}
		pathOfData = filepath.Join(pathOfData, nome)
		fmt.Println(pathOfData)
	} else {
		if !filepath.IsAbs(pathOfProject) {
			pathOfData, err = filepath.Abs(pathOfData)
			if err != nil {
				return err
			}
		}
		pathOfData = filepath.Join(pathOfData, "./.thtml", nome)
	}
	err = file.CreateFolder(pathOfData)
	if err != nil {
		return err
	}

	obj := map[string]any{
		"steps":             map[string]string{},
		"templates":         map[string]string{},
		"folder-all-branch": localAbs,
		"path-of-data":      pathOfData,
	}

	path, err := executable.GetWithJoin("./data/all.json")
	if err != nil {
		return err
	}
	jsonfuncs.InsertInAMap(nome, obj, path)
	return nil
}
