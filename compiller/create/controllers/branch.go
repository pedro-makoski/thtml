package controllers

import (
	"fmt"
	"thtml/compiller/create/estruturas"
	"thtml/funcoes"
)

func CreateBranch(params estruturas.CreateParams, data map[string]any) error {
	err := funcoes.TerminalExec(params.Params["on"].(string), "git", "checkout", "-b", params.Start)
	if err != nil {
		return err
	}

	fmt.Println("Branch criada com sucesso!!")
	return nil
}
