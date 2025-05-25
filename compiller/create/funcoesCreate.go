package create

import (
	"thtml/compiller/create/controllers"
	"thtml/compiller/create/estruturas"
)

func GetCreateFunctions() map[string]estruturas.FuncaoDefinedCreate {
	var CreateFunctions map[string]estruturas.FuncaoDefinedCreate = map[string]estruturas.FuncaoDefinedCreate{
		"branch": {
			HasStart: true,
			Params:   []string{"on"},
			Funcao:   controllers.CreateBranch,
		},
		"folder": {
			HasStart: true,
		},
		"file": {
			HasStart: true,
		},
		"file-on-template": {
			HasStart:  true,
			Params:    []string{"data-base", "with-template-name"},
			Optionals: []string{"change-to-file"},
		},
	}

	CreateFunctions["pasta"] = CreateFunctions["folder"]
	CreateFunctions["arquivo"] = CreateFunctions["file"]
	CreateFunctions["arquivo-com-base-em-template"] = CreateFunctions["file-on-template"]

	return CreateFunctions
}
