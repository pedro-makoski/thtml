package create

type FuncaoDefinedCreate struct {
	HasStart  bool
	Params    []string
	Optionals []string
}

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
