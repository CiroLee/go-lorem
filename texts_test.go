package lorem

import (
	"strings"
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

	r1 := Word(2, "zh")
	r2 := Word(4, "en")

	is.True(unicode.Is(unicode.Han, []rune(r1)[0]))
	is.Equal(len(r1), 6)
	is.True(unicode.Is(unicode.Latin, []rune(r2)[0]))
	is.Equal(len(r2), 4)
}

func TestSentence(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s1 := Sentence(2, "zh")
	s1Arr := strings.Split(s1, "，")
	s2 := Sentence(2, "en")
	s2Arr := strings.Split(s2, " ")

	is.Equal(len(s1Arr), 2)
	is.Equal(len(s2Arr), 2)
}

func TestParagraph(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s1 := Paragraph(2, "zh")
	s1Arr := strings.Split(s1, "。")
	s2 := Paragraph(2, "en")
	s2Arr := strings.Split(s2, ".")

	is.Equal(len(s1Arr), 3)
	is.Equal(len(s2Arr), 3)
}

func TestName(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	en := Name("en", true)
	cn := Name("zh", false)

	is.True(unicode.Is(unicode.Han, []rune(cn)[0]))
	is.True(unicode.Is(unicode.Latin, []rune(en)[0]))
}

func TestStr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s1 := Str(6)
	s2 := Str(0)

	is.Equal(len(s1), 6)
	is.Equal(len(s2), 4)
}
