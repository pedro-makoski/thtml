package utils

func PutValueInMultipleKeys[K comparable](value any, obj map[K]any, keys ...K) map[K]any {
	newObj := CopyMap(obj)

	beforePart := newObj
	var ok bool
	for i, key := range keys {
		if i != len(keys)-1 {
			beforeToBeforePart := beforePart
			beforePart, ok = beforePart[key].(map[K]any)
			if !ok {
				beforeToBeforePart[key] = make(map[K]any)
				beforePart = beforeToBeforePart[key].(map[K]any)
			}

			continue
		}

		beforePart[key] = value
	}

	return newObj
}
