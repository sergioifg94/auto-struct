package generic

import (
	"fmt"
	"reflect"
)

// KeyValueWriter knows how to write a value into a key value collection
type KeyValueWriter struct {
	KeyValueStorage KeyValueStorage

	LevelSeparator   string
	FormatStrategies FormatStrategies
}

// WriteValue writes a given value with a given name into the KeyValueWriter
// key-value collection
func (writer *KeyValueWriter) WriteValue(value interface{}, name string) error {
	return SetValue(value, name, writer.KeyValueStorage, writer.FormatStrategies, writer.LevelSeparator)
}

func SetValue(value interface{}, name string, keyValue KeyValueStorage, formatStrategies FormatStrategies, levelSeparator string) error {
	reflectValue := reflect.ValueOf(value)
	valueKind := reflectValue.Kind()

	if valueKind != reflect.Struct && valueKind != reflect.Slice {
		return fmt.Errorf("Called SetValue function with value of type %s",
			reflectValue.Type().String())
	}

	return SetComplexValue(reflectValue, name, keyValue, formatStrategies, levelSeparator)
}

func SetComplexValue(value reflect.Value, prefix string, keyValue KeyValueStorage, formatStrategies FormatStrategies, levelSeparator string) error {
	primitiveFunctions := map[reflect.Kind]func(reflect.Value) string{
		reflect.String:  func(v reflect.Value) string { return formatStrategies.FormatString(v.String()) },
		reflect.Bool:    func(v reflect.Value) string { return formatStrategies.FormatBool(v.Bool()) },
		reflect.Int64:   func(v reflect.Value) string { return formatStrategies.FormatInt(v.Int()) },
		reflect.Float64: func(v reflect.Value) string { return formatStrategies.FormatFloat(v.Float()) },
	}

	valueKind := value.Kind()
	primitiveFunction, isPrimitive := primitiveFunctions[valueKind]

	// If it's a primitive value directly write it
	if isPrimitive {
		err := keyValue.Set(prefix, primitiveFunction(value))
		if err != nil {
			return err
		}

		return nil
	}

	// If it's a struct, recursively set its fields
	if valueKind == reflect.Struct {
		valueType := value.Type()
		for i := 0; i < valueType.NumField(); i++ {
			fieldType := valueType.Field(i)
			fieldValue := value.Field(i)
			prefixWithField := fmt.Sprintf("%s%s%s", prefix, levelSeparator, fieldType.Name)

			err := SetComplexValue(fieldValue, prefixWithField, keyValue, formatStrategies, levelSeparator)
			if err != nil {
				return err
			}
		}

		return nil
	}

	// If it's a slice, recursively set its elements
	if valueKind == reflect.Slice {
		length := value.Len()

		for i := 0; i < length; i++ {
			prefixWithIndex := fmt.Sprintf("%s%s%d", prefix, levelSeparator, i)
			elem := value.Index(i)
			err := SetComplexValue(elem, prefixWithIndex, keyValue, formatStrategies, levelSeparator)
			if err != nil {
				return err
			}
		}

		return nil
	}

	// If it's a map, recursively set all its fields
	if valueKind == reflect.Map {
		keys := value.MapKeys()

		for _, key := range keys {
			elem := value.MapIndex(key)
			prefixWithKey := fmt.Sprintf("%s%s%s", prefix, levelSeparator, key.String())

			err := SetComplexValue(elem, prefixWithKey, keyValue, formatStrategies, levelSeparator)
			if err != nil {
				return err
			}
		}

		return nil
	}

	return fmt.Errorf("Unsupported type: %s", value.Type().String())
}
