package scanner

import (
	"fmt"
	"os"
	"runtime"
)

type InvalidInputError struct{}

func (e *InvalidInputError) Error() string {
	return "invalid input"
}

func RecoverArgsPanic(group string) {
	if r := recover(); r != nil {
		switch r.(type) {
		case *runtime.TypeAssertionError:
			fmt.Printf("%s: %s", group, &InvalidInputError{})
			os.Exit(1)
		default:
			panic(r)
		}
	}
}

// TODO: check if this is not too slow
type DynamicMap map[string]interface{}

func (m DynamicMap) Map(k string) DynamicMap {
	return m[k].(map[string]interface{})
}

func (m DynamicMap) Val(k string) string {
	return m[k].(string)
}
