package lorem

func Uint(min, max uint) (uint, error) {
	r, err := randomInteger(int(min), int(max))
	if err != nil {
		return 0, err
	}
	return uint(r), nil
}

func Int(min, max int) (int, error) {
	return randomInteger(min, max)
}

func Float64(min, max float64, decimal uint) (float64, error) {
	return randomFloat(min, max, decimal)
}

func Float32(min, max float32, decimal uint) (float32, error) {
	r, err := randomFloat(float64(min), float64(max), decimal)
	if err != nil {
		return 0.0, err
	}
	return float32(r), nil
}
