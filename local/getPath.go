package local

import (
	"os"
	"path/filepath"
)

func GetActualDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

func GetJoinPathInAbsPathActual(relativePath string) (string, error) {
	actualDir, err := GetActualDir()
	if err != nil {
		return "", err
	}
	base := filepath.FromSlash(actualDir)
	joined := filepath.Join(base, relativePath)
	cleaned := filepath.Clean(joined)
	return cleaned, nil
}
