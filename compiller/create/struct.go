package create

type CreateParams struct {
	Name      string
	Params    map[string]any
	Optionals map[string]any
	Configs   FuncaoDefinedCreate
}
