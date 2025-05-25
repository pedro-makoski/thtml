package create

import (
	"thtml/compiller/create/controllers"
	"thtml/compiller/create/estruturas"
)

func GetCreateFunctions() map[string]estruturas.FuncaoDefinedCreate {
	var CreateFunctions map[string]estruturas.FuncaoDefinedCreate = map[string]estruturas.FuncaoDefinedCreate{
		"branch": {
			HasStart:               true,
			Params:                 []string{"on"},
			ParametrosObrigatorios: []int{0},
			Funcao:                 controllers.CreateBranch,
			ValoresSubstituiveis:   []string{"--", "on"},
		},
		"folder": {
			HasStart:             true,
			Funcao:               controllers.CreateFolder,
			ValoresSubstituiveis: []string{"--"},
		},
		"file": {
			Params: []string{"with-text"},
			Defaults: map[int]any{
				0: "",
			},
			Funcao:               controllers.CreateFile,
			HasStart:             true,
			ValoresSubstituiveis: []string{"--", "with-text"},
		},
		"file-on-template": {
			HasStart:             true,
			Params:               []string{"with-template-name"},
			ValoresSubstituiveis: []string{"--", "with-template-name"},
			Funcao:               controllers.CreateFileOnTemplate,
		},
	}

	CreateFunctions["pasta"] = CreateFunctions["folder"]
	CreateFunctions["arquivo"] = CreateFunctions["file"]
	CreateFunctions["arquivo-com-base-em-template"] = CreateFunctions["file-on-template"]

	return CreateFunctions
}
