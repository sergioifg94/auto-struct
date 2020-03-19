package stringmap

import (
	"reflect"
	"testing"
)

func Test_StructFromMap(t *testing.T) {
	reader := NewMapReader(".", map[string]string{
		"testType.Value1":               "hello",
		"testType.InnerValue.Value2":    "12",
		"testType.InnerValue.Value3":    "true",
		"testType.InnerValue.Value4":    "3.14",
		"testType.SliceValue.0":         "foo",
		"testType.SliceValue.1":         "bar",
		"testType.SliceStruct.0.Value2": "45",
		"testType.SliceStruct.0.Value3": "true",
		"testType.SliceStruct.0.Value4": "2.7",
		"testType.SliceStruct.1.Value2": "48",
		"testType.SliceStruct.1.Value3": "true",
		"testType.SliceStruct.1.Value4": "1.52",
		"testType.MapValue.anything":    "15",
		"testType.MapStruct.12.Value2":  "12",
		"testType.MapStruct.12.Value3":  "true",
		"testType.MapStruct.12.Value4":  "5.6",
		"testType.MapStruct.foo.Value2": "2",
		"testType.MapStruct.foo.Value3": "true",
		"testType.MapStruct.foo.Value4": "1.8",
	})

	result := &testType{}
	err := reader.ReadValue("testType", result)

	if err != nil {
		t.Fatal(err)
	}

	expected := testType{
		Value1: "hello",
		InnerValue: innerTestType{
			Value2: 12,
			Value3: true,
			Value4: 3.14,
		},
		SliceValue: []string{"foo", "bar"},
		SliceStruct: []innerTestType{
			innerTestType{
				Value2: 45,
				Value3: true,
				Value4: 2.7,
			},
			innerTestType{
				Value2: 48,
				Value3: true,
				Value4: 1.52,
			},
		},
		MapValue: map[string]int64{
			"anything": 15,
		},
		MapStruct: map[string]innerTestType{
			"12": innerTestType{
				Value2: 12,
				Value3: true,
				Value4: 5.6,
			},
			"foo": innerTestType{
				Value2: 2,
				Value3: true,
				Value4: 1.8,
			},
		},
	}

	if !reflect.DeepEqual(expected, *result) {
		t.Fatalf("Unexpected result. Expected %#v, got %#v", expected, *result)
	}
}
