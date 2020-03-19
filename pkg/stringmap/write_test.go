package stringmap

import (
	"testing"

	"github.com/sergioifg94/auto-struct/pkg/generic"
)

func Test_WriteValue(t *testing.T) {
	expected := testType{
		Value1: "hello",
		InnerValue: innerTestType{
			Value2: 12,
			Value3: true,
		},
		SliceValue: []string{"foo", "bar"},
		SliceStruct: []innerTestType{
			innerTestType{
				Value2: 45,
				Value3: true,
			},
			innerTestType{
				Value2: 48,
				Value3: true,
			},
		},
		MapValue: map[string]int64{
			"anything": 15,
		},
		MapStruct: map[string]innerTestType{
			"12": innerTestType{
				Value2: 12,
				Value3: true,
			},
			"foo": innerTestType{
				Value2: 2,
				Value3: true,
			},
		},
	}

	targetMap := map[string]string{}

	writer := NewMapWriter(&generic.DefaultFormatConvention, ".", targetMap)
	err := writer.WriteValue(expected, "testValue")

	expectedMap := map[string]string{
		"testValue.Value1":               "hello",
		"testValue.InnerValue.Value2":    "12",
		"testValue.InnerValue.Value3":    "true",
		"testValue.SliceValue.0":         "foo",
		"testValue.SliceValue.1":         "bar",
		"testValue.SliceStruct.0.Value2": "45",
		"testValue.SliceStruct.0.Value3": "true",
		"testValue.SliceStruct.1.Value2": "48",
		"testValue.SliceStruct.1.Value3": "true",
		"testValue.MapValue.anything":    "15",
		"testValue.MapStruct.12.Value2":  "12",
		"testValue.MapStruct.12.Value3":  "true",
		"testValue.MapStruct.foo.Value2": "2",
		"testValue.MapStruct.foo.Value3": "true",
	}

	if err != nil {
		t.Fatal(err)
	}

	assertMapEquals(t, expectedMap, targetMap)
}

func assertMapEquals(t *testing.T, expected map[string]string, m map[string]string) {
	for expectedKey, expectedValue := range expected {
		value, ok := m[expectedKey]

		if !ok {
			t.Fatalf("Key %s not found in map", expectedKey)
		}

		if value != expectedValue {
			t.Fatalf("Expected value for key %s to be %s, got %s",
				expectedKey, expectedValue, value)
		}
	}
}
