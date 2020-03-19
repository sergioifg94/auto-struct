package env

import (
	"os"

	"github.com/sergioifg94/auto-struct/pkg/generic"
)

// NewEnvWriter creates a KeyValueWriter that can write values to the environment variables
// using a given levelSeparator and formatStrategies. If the zero value is passed
// for any of these parameters, it's defaulted to "." as levelSeparator, or the
// DefaultFormatConvention as formatStrategies
//
// Note that not all characters are allowed as part of environment variable so
// they can't be used as level separator
func NewEnvWriter(formatStrategies *generic.FormatStrategies, levelSeparator string) *generic.KeyValueWriter {
	if levelSeparator == "" {
		levelSeparator = "_"
	}

	if formatStrategies == nil {
		formatStrategies = &generic.DefaultFormatConvention
	}

	return &generic.KeyValueWriter{
		KeyValueStorage:  *envStorage,
		FormatStrategies: *formatStrategies,
		LevelSeparator:   levelSeparator,
	}
}

var envStorage = &generic.KeyValueStorage{
	Set: os.Setenv,
}
