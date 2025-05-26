package controllers

import (
	"thtml/compiller/create/estruturas"
	"thtml/jsonfuncs"
)

func CreateJsonInsert(params estruturas.CreateParams, data map[string]any) error {
	jsonData, err := jsonfuncs.Dessirealizar[map[string]any](params.Params["data"].(string))
	if err != nil {
		return err
	}

	err = jsonfuncs.InsertObjInJson(jsonData, params.Start)
	if err != nil {
		return err
	}

	err = jsonfuncs.OrderJsonPath(params.Start, params.Params["order-by"].(string), params.Params["order"].(string))

	return err
}
