package commands

import (
	"errors"
	"fmt"
	"os"

	"thtml/controllers/do"
	"thtml/controllers/newsteps"
	"thtml/controllers/newtemplate"
	"thtml/controllers/refresh"
	"thtml/controllers/startproject"
)

func Help() {
	fmt.Println("Insira: init <nome-do-projeto> <local-do-projeto> - para criar um novo projeto")
	fmt.Println("Insira: new-template <local-do-arquivo> <nome-do-projeto> - para inserir templates no projeto")
	fmt.Println("Insira: new-steps <local-do-arquivo> <nome-do-projeto> - para inserir um novo passo a passo no projeto")
	fmt.Println("Insira: do <nome-do-command> <nome-do-projeto> - para fazer um steps no projeto no novo projeto")
	fmt.Println("Insira: refresh - para atualizar todos os nomes")
}

func HandleCommands() error {
	allParams := os.Args[1:]
	if len(allParams) <= 0 {
		fmt.Println("Estão faltando parâmetros")
		Help()
		return errors.New("insira mais parâmetros")
	}

	commandToBranch := allParams[0]
	commandRest := allParams[1:]
	switch commandToBranch {
	case "init":
		err := startproject.InitProject(commandRest)
		if err != nil {
			return err
		}
	case "new-template":
		err := newtemplate.StartTemplate(commandRest)
		if err != nil {
			return err
		}
	case "new-steps":
		err := newsteps.NewSteps(commandRest)
		if err != nil {
			return err
		}
	case "do":
		err := do.Do(commandRest)
		if err != nil {
			return err
		}
	case "refresh":
		err := refresh.RefreshAll()
		if err != nil {
			return err
		}
	default:
		fmt.Println("Você passou um comando inexistente")
		Help()
	}

	return nil
}
