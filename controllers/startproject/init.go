package startproject

import (
	"errors"
	"thtml/executable"
	"thtml/jsonfuncs"
	"thtml/local"
)

func InitProject(params []string) error {
	if len(params) < 1 {
		return errors.New("Coloque o nome do projeto")
	}

	nome := params[0]

	localRelative := "./"
	if len(params) >= 2 {
		localRelative = params[1]
		if localRelative == "." {
			localRelative = "./"
		}
	}
	localAbs, err := local.GetJoinPathInAbsPathActual(localRelative)
	if err != nil {
		return err
	}

	obj := map[string]any{
		"steps":             map[string]string{},
		"templates":         map[string]string{},
		"folder-all-branch": localAbs,
	}

	path, err := executable.GetWithJoin("./data/all.json")
	if err != nil {
		return err
	}
	jsonfuncs.InsertInAMap(nome, obj, path)
	return nil
}
