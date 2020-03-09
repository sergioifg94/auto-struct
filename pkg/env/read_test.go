package env

import (
	"os"
	"testing"
)

type innerTestType struct {
	Value2 int64
	Value3 bool
}

type testType struct {
	Value1     string
	InnerValue innerTestType
}

func Test_StructFromEnv(t *testing.T) {
	result := &testType{}

	os.Setenv("testType_Value1", "hello")
	os.Setenv("testType_InnerValue_Value2", "12")
	os.Setenv("testType_InnerValue_Value3", "true")

	err := StructFromEnv(result, "testType", "_")

	if err != nil {
		t.Fatal(err)
	}

	if result.Value1 != "hello" {
		t.Fatalf("Expected testType.Value1 to be \"hello\", got %s", result.Value1)
	}
	if result.InnerValue.Value2 != 12 {
		t.Fatalf("Expected testType.InnerValue.Value2 to be 12, got %d",
			result.InnerValue.Value2)
	}
	if result.InnerValue.Value3 != true {
		t.Fatalf("Expected testType.InnerValue.Value3 to be true, got %v",
			result.InnerValue.Value3)
	}
}
