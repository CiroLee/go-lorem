package lorem

// return random elements of the array at special num
func Elements[T any](arr []T, num uint) []T {
	var r []T
	for i := 0; i < int(num); i++ {
		r = append(r, randomElement(arr))
	}
	return r
}
