package stringmap

import (
	"reflect"
	"testing"
)

type innerTestType struct {
	Value2 int64
	Value3 bool
}

type testType struct {
	SliceValue  []string
	SliceStruct []innerTestType
	Value1      string
	InnerValue  innerTestType
}

func Test_StructFromMap(t *testing.T) {
	result := &testType{}

	err := StructFromMap(result, "testType", ".", map[string]string{
		"testType.Value1":               "hello",
		"testType.InnerValue.Value2":    "12",
		"testType.InnerValue.Value3":    "true",
		"testType.SliceValue.0":         "foo",
		"testType.SliceValue.1":         "bar",
		"testType.SliceStruct.0.Value2": "45",
		"testType.SliceStruct.0.Value3": "true",
	})

	if err != nil {
		t.Fatal(err)
	}

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
		},
	}

	if !reflect.DeepEqual(expected, *result) {
		t.Fatalf("Unexpected result. Expected %#v, got %#v", expected, *result)
	}
}
