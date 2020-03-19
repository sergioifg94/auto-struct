package stringmap

import (
	"github.com/sergioifg94/auto-struct/pkg/generic"
)

// NewMapWriter creates a KeyValueWriter that can write values to a `map[string]string`
// using a given levelSeparator and formatStrategies. If the zero value is passed
// for any of these parameters, it's defaulted to "." as levelSeparator, or
// the DefaultFormatConvention as formatStrategies
func NewMapWriter(formatStrategies *generic.FormatStrategies, levelSeparator string, target map[string]string) *generic.KeyValueWriter {
	if levelSeparator == "" {
		levelSeparator = "."
	}

	if formatStrategies == nil {
		formatStrategies = &generic.DefaultFormatConvention
	}

	return &generic.KeyValueWriter{
		KeyValueStorage:  mapStorage(target),
		FormatStrategies: *formatStrategies,
		LevelSeparator:   levelSeparator,
	}
}

var mapStorage = func(target map[string]string) generic.KeyValueStorage {
	return generic.KeyValueStorage{
		Set: func(key, value string) error {
			target[key] = value
			return nil
		},
	}
}
