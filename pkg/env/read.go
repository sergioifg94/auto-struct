package env

import (
	"os"
	"regexp"

	"github.com/sergioifg94/auto-struct/pkg/generic"
)

var extractKeyRegex *regexp.Regexp = regexp.MustCompile(`([^=]+)=.*$`)

func StructFromEnv(placeholder interface{}, name, levelSeparator string) error {
	envAccess := generic.KeyValueRetrieval{
		Get: func(key string) (string, bool, error) {
			value := os.Getenv(key)
			value, ok := os.LookupEnv(key)
			return value, ok, nil
		},

		AnyKey: func(predicate func(string) bool) (bool, error) {
			environmentVariables := os.Environ()

			for _, env := range environmentVariables {
				if predicate(extractKey(env)) {
					return true, nil
				}
			}

			return false, nil
		},
	}

	return generic.Struct(placeholder, name, envAccess, levelSeparator)
}

// Extracts the key from a string in the format key=value, as returned in the
// os.Environ function
func extractKey(env string) string {
	matches := extractKeyRegex.FindStringSubmatch(env)
	return matches[1]
}
