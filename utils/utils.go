package utils

import (
	"fmt"
	"strings"
)

// get the length of the decimal part of a number
func GetDecimalPartLen[T float64 | float32](num T) int {
	tmp := strings.Split(fmt.Sprintf("%v", num), ".")
	if len(tmp) <= 1 {
		return 0
	}
	return len(tmp[1])
}

// change the first letter of the string to upper
func Capitalize(s string) string {
	var r []rune
	var rs = []rune(s)
	for i, v := range rs {
		if i == 0 {
			f := strings.ToUpper(string(v))
			r = append(r, []rune(f)[0])
		} else {
			r = append(r, v)
		}
	}

	return string(r)
}
