package stringmap

import "testing"

type innerTestType struct {
	Value2 int64
	Value3 bool
}

type testType struct {
	Value1     string
	InnerValue innerTestType
}

func Test_StructFromMap(t *testing.T) {
	result := &testType{}

	err := StructFromMap(result, "testType", ".", map[string]string{
		"testType.Value1":            "hello",
		"testType.InnerValue.Value2": "12",
		"testType.InnerValue.Value3": "true",
	})

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
