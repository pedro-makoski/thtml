package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFile(path string) (string, error) {
	text, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(text), nil
}

func GetLinesOfString(path string, linesInOrder []int) ([]string, error) {
	lines := make([]string, len(linesInOrder))
	arquivo, err := os.Open(path)
	if err != nil {
		fmt.Printf("Erro ao encontrar o arquivo: %v \n", path)
		return lines, err
	}

	defer arquivo.Close()
	leitor := bufio.NewReader(arquivo)
	i := 0
	posInLine := 0
	for {
		linha, err := leitor.ReadString('\n')
		if err == io.EOF || (posInLine == len(linesInOrder) && len(linesInOrder) != 0) {
			break
		}

		if err != nil {
			return lines, err
		}
		if i == linesInOrder[posInLine] || len(linesInOrder) == 0 {
			linha = strings.TrimSpace(linha)
			lines[posInLine] = linha
			posInLine++
		}

		i++
	}

	return lines, nil
}
