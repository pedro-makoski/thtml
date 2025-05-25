package do

import (
	"errors"
	"fmt"
	"thtml/compiller/interpretmain"
	"thtml/file"
	"thtml/stepsfuncs"
)

func Do(comands []string) error {
	if len(comands) < 2 {
		fmt.Println("Insira: do <nome-do-steps> <nome-do-projeto> - para fazer um steps no projeto no novo projeto")
		err := errors.New("Quantidade de comandos insuficientes")
		return err
	}

	nomeDoSteps := comands[0]
	nomeDoProjeto := comands[1]

	path, err := stepsfuncs.GetStepFullPath(nomeDoProjeto, nomeDoSteps)
	if err != nil {
		return err
	}
	content, err := file.ReadFile(path)
	if err != nil {
		return err
	}

	allBranchPath, err := stepsfuncs.GetAllBranchPath(nomeDoProjeto)
	if err != nil {
		return err
	}
	data := map[string]any{
		"all-branchs":  allBranchPath,
		"project-name": nomeDoProjeto,
	}

	err = interpretmain.Interpret(content, data)
	return err
}
