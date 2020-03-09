package stringmap

import "github.com/sergioifg94/auto-struct/pkg/generic"

// StructFromMap obtains a struct from a map[string]string, unflattens the
// map allowing to obtain hierarchies
func StructFromMap(placeholder interface{}, name, levelSeparator string, source map[string]string) error {
	mapAccess := generic.KeyValueRetrieval{
		Get: func(key string) (string, bool, error) {
			value, ok := source[key]
			return value, ok, nil
		},
	}

	return generic.Struct(placeholder, name, mapAccess, levelSeparator)
}
