package stepsfuncs

import (
	"errors"
	"regexp"
	"thtml/file"
)

func GetComandName(path string) (string, error) {
	lines, err := file.GetLinesOfString(path, []int{0})
	if err != nil {
		return "", err
	}
	firstLine := lines[0]
	re := regexp.MustCompile(`command\s+'?([\S\s]+)'?`)
	matches := re.FindAllStringSubmatch(firstLine, 1)
	for _, match := range matches {
		return match[1], nil
	}

	err = errors.New("Não encontrado nenhum comand na primeira linha. Certifique que esteja na primeira linha")
	return "", err
}
