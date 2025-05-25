package create

type FuncaoDefinedCreate struct {
	HasStart  bool
	Params    []string
	Optionals []string
}

func GetCreateFunctions() map[string]FuncaoDefinedCreate {
	var CreateFunctions map[string]FuncaoDefinedCreate = map[string]FuncaoDefinedCreate{
		"branch": {
			HasStart: true,
			Params:   []string{"on"},
		},
		"folder": {
			HasStart: true,
		},
		"file-on-template": {
			HasStart:  true,
			Params:    []string{"data-base", "with-template-name"},
			Optionals: []string{"change-to-file"},
		},
	}

	CreateFunctions["pasta"] = CreateFunctions["folder"]
	CreateFunctions["arquivo-com-base-em-template"] = CreateFunctions["file-on-template"]

	return CreateFunctions
}
