package file

import "os"

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func InteligentCopy(origem string, destino string) error {
	if IsDir(origem) {
		return CopyDir(origem, destino)
	}

	return CopyFile(origem, destino)

}
