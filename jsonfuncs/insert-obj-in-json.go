package jsonfuncs

func InsertObjInJson(jsonData map[string]any, path string) error {
	if jsonData == nil {
		jsonData = make(map[string]any)
	}

	jsonFull, err := DessirealizarELerArquivo[[]map[string]any](path)
	if err != nil {
		return err
	}

	jsonFull = append(jsonFull, jsonData)

	return SerrializarAndWriteFile(jsonFull, path)
}
