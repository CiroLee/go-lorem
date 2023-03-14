package lorem

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
