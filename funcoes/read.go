package funcoes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Read(inputName string) (string, error) {
	fmt.Print(inputName)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input := strings.TrimSpace(text)
	return input, err
}
