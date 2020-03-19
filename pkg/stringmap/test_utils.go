package stringmap

type innerTestType struct {
	Value2 int64
	Value3 bool
	Value4 float64
}

type testType struct {
	SliceValue  []string
	SliceStruct []innerTestType
	Value1      string
	InnerValue  innerTestType
	MapValue    map[string]int64
	MapStruct   map[string]innerTestType
}
