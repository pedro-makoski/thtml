package utils

func GetIndexOfValue[T comparable](list []T, value T) int {
	for index, item := range list {
		if item == value {
			return index
		}
	}
	return -1
}
