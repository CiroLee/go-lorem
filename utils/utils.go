package utils

import (
	"fmt"
	"strings"
)

// GetDecimalPartLen get the length of the decimal part of a number
func GetDecimalPartLen[T float64 | float32](num T) int {
	tmp := strings.Split(fmt.Sprintf("%v", num), ".")
	if len(tmp) <= 1 {
		return 0
	}
	return len(tmp[1])
}
