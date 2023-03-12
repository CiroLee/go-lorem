package lorem

import "time"

func dateRange(from, to time.Time) time.Time {
	end, _ := randomInteger(int(from.UnixNano()), int(to.UnixNano()))
	return time.Unix(0, int64(end)).UTC()
}

// return a random timestamp support second millisecond and nanosecond via format param
func Timestamp(from, to time.Time, format string) int {
	if format == "nano" {
		return int(dateRange(from, to).UnixNano())
	} else if format == "milli" {
		return int(dateRange(from, to).UnixNano() / 1e3)
	}
	return int(dateRange(from, to).Unix())
}

func Date(from, to time.Time) time.Time {
	return dateRange(from, to)
}
