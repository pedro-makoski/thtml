package controllers

import (
	"thtml/compiller/create/estruturas"
	"thtml/file"
)

func CreateFile(params estruturas.CreateParams, data map[string]any) error {
	return file.CreateFile(params.Start, params.Params["with-text"].(string))
}
