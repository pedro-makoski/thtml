package estruturas

type CreateParams struct {
	Name      string
	Params    map[string]any
	Optionals map[string]any
	Start     string
	Configs   FuncaoDefinedCreate
}

type FuncaoDefinedCreate struct {
	HasStart  bool
	Params    []string
	Optionals []string
	Funcao    func(CreateParams, map[string]any) error
}
