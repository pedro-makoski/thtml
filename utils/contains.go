package utils

func Contains[T comparable](arr []T, value T) bool {
	pos := GetIndexOfValue(arr, value)
	return pos != -1
}
