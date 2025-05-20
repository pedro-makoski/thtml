package file

import "os"

func ThisFileExists(path string) bool {
	_, err := os.Stat(path)

	if(os.IsNotExist(err)) {
		return false
	}

	return true 
}