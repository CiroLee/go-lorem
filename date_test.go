package lorem

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimestamp(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	from := time.Date(2022, time.Month(1), 1, 0, 0, 0, 0, time.Now().Location())
	end := time.Date(2023, time.Month(12), 31, 0, 0, 0, 0, time.Now().Location())
	t1 := Timestamp(from, end, "sec")
	t2 := Timestamp(from, end, "milli")
	t3 := Timestamp(from, end, "nano")

	is.Equal(len(strconv.Itoa(t1)), 10)
	is.Equal(len(strconv.Itoa(t2)), 16)
	is.Equal(len(strconv.Itoa(t3)), 19)
}

func TestDate(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	from := time.Date(2022, time.Month(1), 1, 0, 0, 0, 0, time.Now().Location())
	end := time.Date(2023, time.Month(12), 31, 0, 0, 0, 0, time.Now().Location())
	d := Date(from, end)

	is.True(from.Before(d))
	is.True(end.After(d))
}
