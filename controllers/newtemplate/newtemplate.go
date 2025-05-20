package newtemplate

import (
	"thtml/controllers/configonjsonpart"
	"thtml/templatesfuncs"
)

func StartTemplate(comands []string) error {
	return configonjsonpart.StartSomething(comands, "./data/steps/%v/templates/%v.html", "templates", templatesfuncs.GetDefineNameId)
}