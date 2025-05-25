package controllers

import (
	"thtml/compiller/create/estruturas"
	"thtml/file"
)

func CreateFolder(params estruturas.CreateParams, data map[string]any) error {
	return file.CreateFolder(params.Start)
}
