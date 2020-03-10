package generic

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Struct obtains a struct via reflection given a key-value collection
func Struct(placeholder interface{}, name string, keyValue KeyValueRetrieval, levelSeparator string) error {
	result, err := GetComplexValue(reflect.TypeOf(placeholder).Elem(), name, keyValue, levelSeparator)

	if err != nil {
		return err
	}

	reflect.ValueOf(placeholder).Elem().Set(result)

	return nil
}

// GetComplexValue obtains via reflection a dynamically typed value given is
// expected type, the value key and the logic to obtain that value
func GetComplexValue(outputType reflect.Type, prefix string, keyValue KeyValueRetrieval, levelSeparator string) (reflect.Value, error) {
	primitiveFunctions := map[reflect.Kind]func(string) (interface{}, error){
		reflect.String:  func(x string) (interface{}, error) { return x, nil },
		reflect.Bool:    func(x string) (interface{}, error) { return strconv.ParseBool(x) },
		reflect.Int64:   func(x string) (interface{}, error) { return strconv.ParseInt(x, 10, 64) },
		reflect.Float64: func(x string) (interface{}, error) { return strconv.ParseFloat(x, 64) },
	}

	outputKind := outputType.Kind()
	primitiveFunc, isPrimitive := primitiveFunctions[outputKind]

	// If it's a primitive value, directly handle it
	if isPrimitive {
		value, ok, err := keyValue.Get(prefix)

		if err != nil {
			return reflect.Zero(outputType), fmt.Errorf("Error getting value for key %s", prefix)
		}
		if !ok {
			return reflect.Zero(outputType), fmt.Errorf("Value not found for %s", prefix)
		}

		result, err := primitiveFunc(value)
		if err != nil {
			return reflect.Zero(outputType), err
		}

		return reflect.ValueOf(result), nil
	}

	// If it's a struct, recursively get its fields
	if outputKind == reflect.Struct {
		result := reflect.New(outputType)

		for i := 0; i < outputType.NumField(); i++ {
			field := outputType.Field(i)
			fieldValue, err := GetComplexValue(field.Type, fmt.Sprintf("%s%s%s", prefix, levelSeparator, field.Name), keyValue, levelSeparator)

			if err != nil {
				return reflect.Zero(outputType), err
			}

			result.Elem().Field(i).Set(fieldValue)
		}

		return result.Elem(), nil
	}

	// If it's a struct, recursively get its elements
	if outputKind == reflect.Slice {
		innerType := outputType.Elem()
		result := reflect.MakeSlice(outputType, 0, 0)

		// Look for the keys that start with <PREFIX><LEVEL_SEPARATOR><INDEX>
		// and, if they're found, recursively get the value and add it to the
		// resulting slice
		for i := 0; ; i++ {
			keyPrefix := fmt.Sprintf("%s%s%d", prefix, levelSeparator, i)
			hasIndexKey, err := keyValue.AnyKey(func(key string) bool {
				return strings.HasPrefix(key, keyPrefix)
			})
			if err != nil {
				return reflect.Zero(outputType), err
			}

			// If no key is found, assume that there's no more elements in the
			// slice
			if !hasIndexKey {
				break
			}

			elemValue, err := GetComplexValue(innerType, keyPrefix, keyValue, levelSeparator)
			if err != nil {
				return reflect.Zero(outputType), err
			}

			result = reflect.Append(result, elemValue)
		}

		return result, nil
	}

	return reflect.Zero(outputType), fmt.Errorf("Unsupported type: %s", outputType.String())
}
