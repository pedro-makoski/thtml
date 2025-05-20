package executable

import (
	"os"
	"path"
	"path/filepath"
)

func GetPath() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err
	}

	return path, nil
}

func GetFolder() (string, error) {
	pathAct, err := GetPath()
	if err != nil {
		return "", err
	}
	return path.Dir(pathAct), nil
}

func GetWithJoin(actual string) (string, error) {
	actPath, err := GetFolder()
	if err != nil {
		return "", err
	}
	joined := filepath.Join(actPath, actual)
	cleaned := filepath.Clean(joined)
	return cleaned, nil
}
