package templatesfuncs

import (
	"thtml/controllers/configonjsonpart"
	"thtml/file"
)

func GetPath(projectName string, templateName string) (string, error) {
	return configonjsonpart.GetPathOfJSON(templateName, projectName, "templates")
}

func GetTextOfTemplate(projectName string, templateName string) (string, error) {
	path, err := GetPath(projectName, templateName)
	if err != nil {
		return "", err
	}

	return file.ReadFile(path)
}
