package interpretmain

import (
	"fmt"
	"thtml/compiller/funcoes"
	"thtml/utils"
)

func Interpret(file string, presetData map[string]any) error {
	data := presetData

	var AllMethods = map[string](func(string, map[string]any) error){
		"ask":                  funcoes.Ask,
		"ask:":                 funcoes.Ask,
		"perguntar":            funcoes.Ask,
		"perguntar:":           funcoes.Ask,
		"say":                  funcoes.Say,
		"say:":                 funcoes.Say,
		"falar":                funcoes.Say,
		"falar:":               funcoes.Say,
		"define":               funcoes.Define,
		"definir":              funcoes.Define,
		"copy":                 funcoes.Copy,
		"copiar":               funcoes.Copy,
		"terminal-exec":        funcoes.TerminalExec,
		"executar-no-terminal": funcoes.TerminalExec,
		"terminal":             funcoes.TerminalExec,
		"create":               funcoes.Create,
		"read":                 funcoes.Read,
		"ler":                  funcoes.Read,
		"remove":               funcoes.Remove,
		"remover":              funcoes.Remove,
		"define-as-file-name":  funcoes.DefineAsFile,
		"//":                   func(line string, data map[string]any) error { return nil },
	}

	lines := utils.TransformStringInList(file)

	lines = lines[1:]

	for _, line := range lines {
		comands := utils.SplitBy(line, []string{`\s`})
		mainComand := comands[0]
		funcao, ok := AllMethods[mainComand]
		if !ok {
			return fmt.Errorf("comando %s não existe", mainComand)
		}
		err := funcao(line[len([]byte(mainComand))+1:], data)
		if err != nil {
			return err
		}
	}

	return nil
}
