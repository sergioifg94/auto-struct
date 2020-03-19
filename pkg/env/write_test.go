package env

import (
	"os"
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

	writer := NewEnvWriter(&generic.DefaultFormatConvention, "_")
	err := writer.WriteValue(expected, "testValue")

	expectedEnv := map[string]string{
		"testValue_Value1":               "hello",
		"testValue_InnerValue_Value2":    "12",
		"testValue_InnerValue_Value3":    "true",
		"testValue_SliceValue_0":         "foo",
		"testValue_SliceValue_1":         "bar",
		"testValue_SliceStruct_0_Value2": "45",
		"testValue_SliceStruct_0_Value3": "true",
		"testValue_SliceStruct_1_Value2": "48",
		"testValue_SliceStruct_1_Value3": "true",
		"testValue_MapValue_anything":    "15",
		"testValue_MapStruct_12_Value2":  "12",
		"testValue_MapStruct_12_Value3":  "true",
		"testValue_MapStruct_foo_Value2": "2",
		"testValue_MapStruct_foo_Value3": "true",
	}

	if err != nil {
		t.Fatal(err)
	}

	for env, expectedValue := range expectedEnv {
		value, ok := os.LookupEnv(env)

		if !ok {
			t.Fatalf("Expected environment variable %s to be created, but wasn't found",
				env)
		}

		if value != expectedValue {
			t.Fatalf("Expected environment variable %s to be %s but got %s",
				env, expectedValue, value)
		}
	}
}
