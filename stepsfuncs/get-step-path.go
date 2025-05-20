package stepsfuncs

import (
	"thtml/controllers/configonjsonpart"
	"thtml/executable"
	"thtml/jsonfuncs"
)

func GetStepFullPath(project string, stepName string) (string, error) {
	return configonjsonpart.GetPathOfJSON(stepName, project, "steps")
}

func GetAllBranchPath(project string) (string, error) {
	allJson, err := executable.GetWithJoin("./data/all.json")
	if err != nil {
		return "", err
	}
	return jsonfuncs.GetValueOfKey[string](allJson, project, "folder-all-branch")
}
