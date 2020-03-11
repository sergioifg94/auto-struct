package generic

// KeyValueRetrieval is an abstraction of the strategy to access a key value collection
type KeyValueRetrieval struct {

	// Get retrieves a value given a string. If the value doesn't exist,
	// returns an empty string and false as the second result. If there
	// was an error retrieving the value, return the error as the third
	// value
	Get func(key string) (string, bool, error)

	// AnyKey checks if there's at least one key in the collection that
	// satisfies a given predicate
	AnyKey func(predicate func(string) bool) (bool, error)

	// FilterKeys creates a slice of keys that satisfy a given predicate
	FilterKeys func(predicate func(string) bool) ([]string, error)
}
