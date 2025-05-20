package funcoes

import (
	"fmt"
	"regexp"
)

func ReplaceProperties(text string, data map[string]any) string {
	newText := text
	for key, value := range data {
		re := regexp.MustCompile(fmt.Sprintf(`\{\{((?:\s*)%s(?:\s*))\}\}`, key))
		newText = string(re.ReplaceAll([]byte(newText), []byte(fmt.Sprintf("%v", value))))
	}

	return newText
}
