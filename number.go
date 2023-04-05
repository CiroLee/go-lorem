package lorem

import (
	"fmt"
	"strconv"
)

// return a random unit
func Uint(min, max uint) (uint, error) {
	r, err := randomInteger(int(min), int(max))
	if err != nil {
		return 0, err
	}
	return uint(r), nil
}

// return a random int
func Int(min, max int) (int, error) {
	return randomInteger(min, max)
}

// return a random int via special digit
func IntBy(digit uint, positive bool) (int, error) {
	var num string
	var prefix = ""
	if !positive {
		prefix = "-"
	}
	if digit == 0 {
		return 0, fmt.Errorf("digit must be greater than zero")
	}
	for i := 0; i < int(digit); i++ {
		if i == 0 && digit == 1 {
			random, _ := randomInteger(0, 9)
			num = strconv.Itoa(random)
		} else if i == 0 && digit != 1 {
			random, _ := randomInteger(1, 9)
			num = strconv.Itoa(random)
		} else {
			random, _ := randomInteger(0, 9)
			num += strconv.Itoa(random)
		}
	}
	return strconv.Atoi(prefix + num)
}

// return a random float32 number between min and max.And include max
func Float32(min, max float32, decimal uint) (float32, error) {
	r, err := randomFloat(float64(min), float64(max), decimal)
	if err != nil {
		return 0.0, err
	}
	return float32(r), nil
}

// return a random float64 number between min and max.Note: not include max
func Float64(min, max float64, decimal uint) (float64, error) {
	return randomFloat(min, max, decimal)
}
