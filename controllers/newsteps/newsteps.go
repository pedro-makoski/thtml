package newsteps

import (
	"thtml/controllers/configonjsonpart"
	"thtml/stepsfuncs"
)

func NewSteps(comands []string) error {
	return configonjsonpart.StartSomething(comands, "./data/steps/%v/%v.sthml", "steps", stepsfuncs.GetComandName)
}