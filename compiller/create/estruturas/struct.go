package estruturas

type CreateParams struct {
	Name      string
	Params    map[string]any
	Optionals map[string]any
	Start     string
	Configs   FuncaoDefinedCreate
}

type FuncaoDefinedCreate struct {
	HasStart               bool
	Params                 []string
	ParametrosObrigatorios []int
	Optionals              []string
	Defaults               map[int]any
	Funcao                 func(CreateParams, map[string]any) error
}
