package lorem

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// generate random integer
func randomInteger(min, max int) (int, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	if max <= min {
		return 0, fmt.Errorf("min must be less than max. min: %d, max: %d", min, max)
	}
	return rand.Intn(max-min) + min, nil
}

// generate random float64 number
func randomFloat(min, max float64, decimal uint) (float64, error) {
	if max <= min {
		return 0.0, fmt.Errorf("min must be less than max. min: %v, max: %v", min, max)
	}
	d := math.Pow10(4)
	if decimal > 0 {
		d = math.Pow10(int(decimal))
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	random := rand.Float64() * (max - min)
	temp := strconv.FormatFloat(math.Trunc(random*d)/d, 'f', -1, 64)
	num, _ := strconv.ParseFloat(temp, 64)
	return num, nil
}
