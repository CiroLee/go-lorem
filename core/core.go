package core

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// RandomInteger generate random integer
func RandomInteger(min, max int) (int, error) {
	rand.Seed(time.Now().UnixNano())
	if max <= min {
		return 0, fmt.Errorf("min must be less than max. min: %d, max: %d", min, max)
	}
	return rand.Intn(max-min) + min, nil
}

// RandomFloat generate random float64 number
func RandomFloat(min, max int, decimal int) (float64, error) {
	if max <= min {
		return 0.0, fmt.Errorf("min must be less than max. min: %d, max: %d", min, max)
	}
	d := math.Pow10(4)
	if decimal > 0 {
		d = math.Pow10(decimal)
	}
	rand.Seed(time.Now().UnixNano())
	random := rand.Float64() * (float64(max) - float64(min))
	temp := strconv.FormatFloat(math.Trunc(random*d)/d, 'f', -1, 64)
	num, _ := strconv.ParseFloat(temp, 64)
	return num, nil
}
