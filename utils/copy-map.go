package utils

import "maps"

func CopyMap[K comparable, V any](obj map[K]V) map[K]V {
	newMap := make(map[K]V)
	maps.Copy(newMap, obj)

	return newMap
}