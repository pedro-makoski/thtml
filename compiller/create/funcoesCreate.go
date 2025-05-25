package create

type FuncaoDefinedCreate struct {
	HasStart  bool
	Params    []string
	Optionals []string
	Funcao    func(CreateParams, map[string]any) error
}

func GetCreateFunctions() map[string]FuncaoDefinedCreate {
	var CreateFunctions map[string]FuncaoDefinedCreate = map[string]FuncaoDefinedCreate{
		"branch": {
			HasStart: true,
			Params:   []string{"on"},
		},
		"folder": {
			HasStart: true,
			Params:   []string{"on"},
		},
		"file": {
			HasStart: true,
			Params:   []string{"on"},
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
