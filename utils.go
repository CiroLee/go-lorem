package lorem

import (
	"fmt"
	"strings"
)

// get the length of the decimal part of a number
func getDecimalPartLen[T float64 | float32](num T) int {
	tmp := strings.Split(fmt.Sprintf("%v", num), ".")
	if len(tmp) <= 1 {
		return 0
	}
	return len(tmp[1])
}
