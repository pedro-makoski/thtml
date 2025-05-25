package newsteps

import (
	"thtml/controllers/configonjsonpart"
	"thtml/stepsfuncs"
)

func NewSteps(comands []string) error {
	return configonjsonpart.StartSomething(comands, "./%v.sthtml", "steps", stepsfuncs.GetComandName)
}
