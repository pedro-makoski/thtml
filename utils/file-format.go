package utils

import (
	"path/filepath"
	"strings"
)

func RemoverAcentos(text string) string {
	replacer := strings.NewReplacer(
		"á", "a", "à", "a", "ã", "a", "â", "a",
		"é", "e", "è", "e", "ê", "e",
		"í", "i", "ì", "i", "î", "i",
		"ó", "o", "ò", "o", "õ", "o", "ô", "o",
		"ú", "u", "ù", "u", "û", "u",
		"ç", "c",
	)
	return replacer.Replace(text)
}

func FileFormat(text string) string {
	ext := filepath.Ext(text)
	text = text[:len(text)-len(ext)]
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, " ", "-")
	text = strings.ReplaceAll(text, "/", "-")
	text = strings.ReplaceAll(text, "?", "")
	text = strings.ReplaceAll(text, "!", "")
	text = strings.ReplaceAll(text, "", "")
	text = RemoverAcentos(text)

	return text + ext
}
