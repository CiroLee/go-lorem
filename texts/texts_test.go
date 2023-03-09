package lorem

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestLetter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	enResult := Letter("en")
	zhResult := Letter("zh")

	is.True(unicode.Is(unicode.Latin, []rune(enResult)[0]))
	is.True(unicode.Is(unicode.Han, []rune(zhResult)[0]))
}

func TestWord(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Word(WordProps{})
	r2 := Word(WordProps{Num: 4, Lang: "en", Upper: true})

	is.True(unicode.Is(unicode.Han, []rune(r1)[0]))
	is.Equal(len(r1), 6)
	is.True(unicode.Is(unicode.Latin, []rune(r2)[0]))
	is.Equal(len(r2), 4)
}
