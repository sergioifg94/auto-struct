package generic

import (
	"fmt"
	"reflect"
	"strconv"
)

// "unflattens" the collection allowing to obtain hierarchies
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
	primitiveFunctions := map[string]func(string) (interface{}, error){
		"string": func(x string) (interface{}, error) { return x, nil },
		"bool":   func(x string) (interface{}, error) { return strconv.ParseBool(x) },
		"int64":  func(x string) (interface{}, error) { return strconv.ParseInt(x, 10, 64) },
		"float":  func(x string) (interface{}, error) { return strconv.ParseFloat(x, 64) },
	}

	primitiveFunc, isHandled := primitiveFunctions[outputType.String()]

	if isHandled {
		value, ok, err := keyValue.Get(prefix)

		if err != nil {
			return reflect.Value{}, fmt.Errorf("Error getting value for key %s", prefix)
		}
		if !ok {
			return reflect.Value{}, fmt.Errorf("Value not found for %s", prefix)
		}

		result, err := primitiveFunc(value)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(result), nil
	}

	if outputType.NumField() == 0 {
		return reflect.Zero(outputType), nil
	}

	result := reflect.New(outputType)
	for i := 0; i < outputType.NumField(); i++ {
		field := outputType.Field(i)
		fieldValue, err := GetComplexValue(field.Type, fmt.Sprintf("%s%s%s", prefix, levelSeparator, field.Name), keyValue, levelSeparator)

		if err != nil {
			return reflect.Value{}, err
		}

		result.Elem().Field(i).Set(fieldValue)
	}

	return result.Elem(), nil
}
