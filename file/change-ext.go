package file

import (
	"path/filepath"
)

func ChangeExt(fullPath string, newExt string) string {
	ext := filepath.Ext(fullPath)
	newPath := fullPath[:len(fullPath)-len(ext)]
	if newExt[0] != '.' {
		newExt = "." + newExt
	}
	newPath += newExt
	return newPath
}
