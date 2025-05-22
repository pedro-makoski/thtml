package funcoes

import (
	"os"
	"os/exec"
)

func TerminalExec(path string, nome string, args ...string) error {
	cmd := exec.Command(nome, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = path

	err := cmd.Run()
	return err
}
