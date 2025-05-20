package utils

import (
	"regexp"
	"strings"
)

func TransformStringInList(text string) []string {
	re := regexp.MustCompile("[\n;]+")
	partes := re.Split(text, -1)
	for i := int(0); i < len(partes); i++ {
		if strings.TrimSpace(partes[i]) == "" {
			partes = DeleteInList(partes, i)
		}
	}
	return partes
}

func SplitBy(text string, possibleValues []string) []string {
	re := regexp.MustCompile(strings.Join(possibleValues, "|"))
	partes := re.Split(text, -1)
	for i := int(0); i < len(partes); i++ {
		if strings.TrimSpace(partes[i]) == "" {
			partes = DeleteInList(partes, i)
		}
	}
	return partes
}
