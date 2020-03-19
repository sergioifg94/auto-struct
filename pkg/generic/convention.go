package generic

import (
	"fmt"
	"strconv"
)

// FormatStrategies is used to delegate the logic to format primitive
// values
type FormatStrategies struct {
	FormatString func(string) string
	FormatBool   func(bool) string
	FormatInt    func(int64) string
	FormatFloat  func(float64) string
}

// DefaultFormatConvention contains default logic to format primitive values
// * FormatString is the identity function
// * FormatBool uses `strconv.FormatBool`
// * FormatFloat uses the %f string format
// * FormatInt uses the %d string format
var DefaultFormatConvention = FormatStrategies{
	FormatString: func(s string) string { return s },
	FormatBool:   strconv.FormatBool,
	FormatFloat: func(f float64) string {
		return fmt.Sprintf("%f", f)
	},
	FormatInt: func(n int64) string {
		return fmt.Sprintf("%d", n)
	},
}
