package env

import (
	"os"

	"github.com/sergioifg94/auto-struct/pkg/generic"
)

func StructFromEnv(placeholder interface{}, name, levelSeparator string) error {
	envAccess := generic.KeyValueRetrieval{
		Get: func(key string) (string, bool, error) {
			value := os.Getenv(key)
			return value, value != "", nil
		},
	}

	return generic.Struct(placeholder, name, envAccess, levelSeparator)
}
