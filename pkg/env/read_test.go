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
}

func Test_StructFromEnv(t *testing.T) {
	result := &testType{}

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

	err := StructFromEnv(result, "testType", "_")

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
