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
		"image-file-from-svg": {
			HasStart:               true,
			Params:                 []string{"color", "width", "height", "on", "format", "stroke"},
			ParametrosObrigatorios: []int{3, 4},
			Defaults: map[int]any{
				0: "",
				1: 0,
				2: 0,
				5: "",
			},
			ValoresSubstituiveis: []string{"--", "color", "width", "height", "on", "format", "stroke"},
			Funcao:               controllers.CreateImageFileFromSvg,
		},
		"json-insert": {
			HasStart:               true,
			Params:                 []string{"data", "order", "order-by"},
			ParametrosObrigatorios: []int{0},
			Defaults: map[int]any{
				1: "none",
				2: "none",
			},
			ValoresSubstituiveis: []string{"--", "data", "order", "order-by"},
			Funcao:               controllers.CreateJsonInsert,
		},
	}

	CreateFunctions["pasta"] = CreateFunctions["folder"]
	CreateFunctions["arquivo"] = CreateFunctions["file"]
	CreateFunctions["arquivo-com-base-em-template"] = CreateFunctions["file-on-template"]

	return CreateFunctions
}
