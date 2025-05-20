package utils

func GetKeysAndValues[K comparable, V any](obj map[K]V) ([]K, []V) {
	keys := []K{}
	values := []V{}

	for key, value := range obj {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}