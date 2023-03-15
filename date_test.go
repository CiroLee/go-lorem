package lorem

import (
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode"

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

	is.Len(strconv.Itoa(t1), 10)
	is.Len(strconv.Itoa(t2), 16)
	is.Len(strconv.Itoa(t3), 19)
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

func TestWeek(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	w1 := Week("en", false)
	w2 := Week("en", true)
	w3 := Week("zh", false)

	is.True(unicode.Is(unicode.Latin, []rune(w1)[0]))
	is.True(strings.Contains(w2, "."))
	is.True(unicode.Is(unicode.Han, []rune(w3)[0]))
}

func TestMonth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	w1 := Month("en", false)
	w2 := Month("en", true)
	w3 := Month("zh", false)

	is.True(unicode.Is(unicode.Latin, []rune(w1)[0]))
	is.True(strings.Contains(w2, "."))
	is.True(unicode.Is(unicode.Han, []rune(w3)[0]))
}

func TestFormatDate(t *testing.T) {
	is := assert.New(t)

	from := time.Date(2022, time.Month(1), 1, 0, 0, 0, 0, time.Now().Location())
	to := time.Date(2023, time.Month(12), 31, 0, 0, 0, 0, time.Now().Location())

	fd := FormatDate(from, to, SlashLayout)
	is.True(strings.Contains(fd, "/"))
}
