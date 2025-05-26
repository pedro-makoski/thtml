package jsonfuncs

import (
	"fmt"
	"reflect"
	"sort"
	"thtml/utils"
)

func OrderJson(orderby, ordertype string, jsonArr []map[string]any) ([]map[string]any, error) {
	if orderby == "" || orderby == "none" || ordertype == "none" || ordertype == "" {
		return jsonArr, nil
	}

	if len(jsonArr) == 0 {
		return jsonArr, nil
	}

	for i, item := range jsonArr {
		if _, ok := item[orderby]; !ok {
			return jsonArr, fmt.Errorf("chave %q não encontrada no elemento %d", orderby, i)
		}
	}

	firstVal := jsonArr[0][orderby]
	kind := reflect.TypeOf(firstVal).Kind()
	isString := kind == reflect.String
	isNumber := kind == reflect.Int ||
		kind == reflect.Int8 ||
		kind == reflect.Int16 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 ||
		kind == reflect.Float32 ||
		kind == reflect.Float64

	if !isString && !isNumber {
		return jsonArr, fmt.Errorf("unsupported type %s for key %q", kind, orderby)
	}

	asc, err := parseOrderType(ordertype, isString)
	if err != nil {
		return jsonArr, err
	}

	sort.Slice(jsonArr, func(i, j int) bool {
		vi := jsonArr[i][orderby]
		vj := jsonArr[j][orderby]

		if isString {
			si := vi.(string)
			sj := vj.(string)
			if asc {
				return si < sj
			}
			return si > sj
		}

		fi := utils.ToFloat64(vi)
		fj := utils.ToFloat64(vj)
		if asc {
			return fi < fj
		}
		return fi > fj
	})

	fmt.Println("Ordenado com sucesso!!!")
	return jsonArr, nil
}

func parseOrderType(ordertype string, isString bool) (bool, error) {
	switch ordertype {
	case "A-Z":
		if isString {
			return true, nil
		}
	case "Z-A":
		if isString {
			return false, nil
		}
	case "0-9":
		if !isString {
			return true, nil
		}
	case "9-0":
		if !isString {
			return false, nil
		}
	}
	return false, fmt.Errorf("Tipo de ordem invalido %q para os valores:", ordertype, map[bool]string{true: "string", false: "numeric"}[isString])
}

func OrderJsonPath(path string, orderby string, ordertype string) error {
	json, err := DessirealizarELerArquivo[[]map[string]any](path)
	if err != nil {
		return err
	}

	json, err = OrderJson(orderby, ordertype, json)
	if err != nil {
		return err
	}

	err = SerrializarAndWriteFile(json, path)

	return err
}
