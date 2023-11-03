package scanner

import (
	"fmt"
)

func ConvertInputToMap(i interface{}) (map[string]string, error) {
	input := make(map[string]string)
	data, ok := i.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input")
	}

	for k, iv := range data {
		v, ok := iv.(string)
		if !ok {
			return nil, fmt.Errorf("invalid input")
		}
		input[k] = v
	}

	return input, nil
}

// TODO: check if this is not too slow
type DynamicMap map[string]interface{}

func (m DynamicMap) Map(k string) DynamicMap {
	return m[k].(map[string]interface{})
}

func (m DynamicMap) Val(k string) string {
	return m[k].(string)
}
