package newtemplate

import (
	"thtml/controllers/configonjsonpart"
	"thtml/templatesfuncs"
)

func StartTemplate(comands []string) error {
	return configonjsonpart.StartSomething(comands, "%v/templates/%v.html", "templates", templatesfuncs.GetDefineNameId)
}
