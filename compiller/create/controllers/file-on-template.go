package controllers

import (
	"errors"
	"html/template"
	"strings"
	"thtml/compiller/create/estruturas"
	"thtml/templatesfuncs"
)

func CreateFileOnTemplate(params estruturas.CreateParams, data map[string]any) error {
	templateName := params.Params["with-template-name"].(string)
	templateText, err := templatesfuncs.GetTextOfTemplate(data["project-name"].(string), templateName)
	if err != nil {
		return err
	}

	htmlTemplate := strings.Split(templateText, "\n")
	if err != nil {
		return err
	}
	toTemplate := strings.Join(htmlTemplate[1:len(htmlTemplate)-2], "\n")
	template, err := template.New(templateName).Parse(toTemplate)
	if err != nil {
		return errors.New("Erro ao criar o template")
	}

	return templatesfuncs.CreateTemplateFile(data, template, params.Start)
}
