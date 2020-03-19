package stringmap

import (
	"github.com/sergioifg94/auto-struct/pkg/generic"
)

// NewMapReader creates a KeyValueReader that reads values from a `map[string]string`
// using a given level separator. If the empty string is passed as levelSeparator,
// defaults to "."
func NewMapReader(levelSeparator string, source map[string]string) *generic.KeyValueReader {
	if levelSeparator == "" {
		levelSeparator = "."
	}

	mapAccess := generic.KeyValueRetrieval{
		Get: func(key string) (string, bool, error) {
			value, ok := source[key]
			return value, ok, nil
		},

		AnyKey: func(predicate func(string) bool) (bool, error) {
			for key := range source {
				if predicate(key) {
					return true, nil
				}
			}

			return false, nil
		},

		FilterKeys: func(predicate func(string) bool) ([]string, error) {
			result := []string{}

			for key := range source {
				if predicate(key) {
					result = append(result, key)
				}
			}

			return result, nil
		},
	}

	return &generic.KeyValueReader{
		KeyValue:       mapAccess,
		LevelSeparator: levelSeparator,
	}
}
