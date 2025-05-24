package startproject

import (
	"errors"
	"path/filepath"
	"thtml/executable"
	"thtml/file"
	"thtml/jsonfuncs"
	"thtml/local"
)

func InitProject(params []string) error {
	if len(params) < 2 {
		return errors.New("Coloque o nome do projeto e o local do projeto(-- para ficar no local do exe)")
	}

	nome := params[0]
	pathOfProject := params[1]

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

	if pathOfProject == "--" {
		pathOfProject, err = executable.GetWithJoin("./data/steps")
		if err != nil {
			return err
		}
	} else {
		if !filepath.IsAbs(pathOfProject) {
			pathOfProject, err = filepath.Abs(pathOfProject)
			if err != nil {
				return err
			}
		}
		pathOfProject = filepath.Join(pathOfProject, "./.thtml")
	}
	file.CreateFolder(pathOfProject)

	obj := map[string]any{
		"steps":             map[string]string{},
		"templates":         map[string]string{},
		"folder-all-branch": localAbs,
		"path-of-data":      pathOfProject,
	}

	path, err := executable.GetWithJoin("./data/all.json")
	if err != nil {
		return err
	}
	jsonfuncs.InsertInAMap(nome, obj, path)
	return nil
}
