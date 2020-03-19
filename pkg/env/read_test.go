package env

import (
	"os"
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

func Test_StructFromEnv(t *testing.T) {
	os.Setenv("testType_Value1", "hello")
	os.Setenv("testType_InnerValue_Value2", "12")
	os.Setenv("testType_InnerValue_Value3", "true")
	os.Setenv("testType_Value1", "hello")
	os.Setenv("testType_InnerValue_Value2", "12")
	os.Setenv("testType_InnerValue_Value3", "true")
	os.Setenv("testType_SliceValue_0", "foo")
	os.Setenv("testType_SliceValue_1", "bar")
	os.Setenv("testType_SliceStruct_0_Value2", "45")
	os.Setenv("testType_SliceStruct_0_Value3", "true")
	os.Setenv("testType_SliceStruct_1_Value2", "48")
	os.Setenv("testType_SliceStruct_1_Value3", "true")
	os.Setenv("testType_MapValue_anything", "15")
	os.Setenv("testType_MapStruct_12_Value2", "12")
	os.Setenv("testType_MapStruct_12_Value3", "true")
	os.Setenv("testType_MapStruct_foo_Value2", "2")
	os.Setenv("testType_MapStruct_foo_Value3", "true")

	result := &testType{}

	reader := NewEnvReader("_")
	err := reader.ReadValue("testType", result)

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
