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
	MapValue    map[string]int64
	MapStruct   map[string]innerTestType
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
		"testType.SliceStruct.1.Value2": "48",
		"testType.SliceStruct.1.Value3": "true",
		"testType.MapValue.anything":    "15",
		"testType.MapStruct.12.Value2":  "12",
		"testType.MapStruct.12.Value3":  "true",
		"testType.MapStruct.foo.Value2": "2",
		"testType.MapStruct.foo.Value3": "true",
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

	if !reflect.DeepEqual(expected, *result) {
		t.Fatalf("Unexpected result. Expected %#v, got %#v", expected, *result)
	}
}
