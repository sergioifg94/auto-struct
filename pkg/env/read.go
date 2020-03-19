package env

import (
	"os"
	"regexp"

	"github.com/sergioifg94/auto-struct/pkg/generic"
)

var extractKeyRegex *regexp.Regexp = regexp.MustCompile(`([^=]+)=.*$`)

// NewEnvReader creates a KeyValueReader that reads values from the environment
// variables using a given level separator. If the empty string is passed as
// levelSeparator, defaults to "."
//
// Note that not all characters are allowed as part of environment variable so
// they can't be used as level separator
func NewEnvReader(levelSeparator string) *generic.KeyValueReader {
	if levelSeparator == "" {
		levelSeparator = "_"
	}

	// Lazy initialization of the environment variables list
	var _environmentVariables []string
	environmentVariables := func() []string {
		if _environmentVariables == nil {
			_environmentVariables = os.Environ()
		}

		return _environmentVariables
	}

	envAccess := generic.KeyValueRetrieval{
		Get: func(key string) (string, bool, error) {
			value := os.Getenv(key)
			value, ok := os.LookupEnv(key)
			return value, ok, nil
		},

		AnyKey: func(predicate func(string) bool) (bool, error) {
			for _, env := range environmentVariables() {
				if predicate(extractKey(env)) {
					return true, nil
				}
			}

			return false, nil
		},

		FilterKeys: func(predicate func(string) bool) ([]string, error) {
			result := []string{}

			for _, env := range environmentVariables() {
				key := extractKey(env)
				if predicate(key) {
					result = append(result, key)
				}
			}

			return result, nil
		},
	}

	return &generic.KeyValueReader{
		KeyValue:       envAccess,
		LevelSeparator: levelSeparator,
	}
}

// Extracts the key from a string in the format key=value, as returned in the
// os.Environ function
func extractKey(env string) string {
	matches := extractKeyRegex.FindStringSubmatch(env)
	return matches[1]
}
