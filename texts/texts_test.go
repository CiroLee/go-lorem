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

	r1 := Word(BaseProps{})
	r2 := Word(BaseProps{Num: 4, Lang: "en"})

	is.True(unicode.Is(unicode.Han, []rune(r1)[0]))
	is.Equal(len(r1), 6)
	is.True(unicode.Is(unicode.Latin, []rune(r2)[0]))
	is.Equal(len(r2), 4)
}

func TestSentence(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s1 := Sentence(BaseProps{Num: 2})
	s1Arr := strings.Split(s1, "，")
	s2 := Sentence(BaseProps{Num: 2, Lang: "en"})
	s2Arr := strings.Split(s2, " ")

	is.Equal(len(s1Arr), 2)
	is.Equal(len(s2Arr), 2)
}

func TestParagraph(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	s1 := Paragraph(BaseProps{Num: 2})
	s1Arr := strings.Split(s1, "。")
	s2 := Paragraph(BaseProps{Num: 2, Lang: "en"})
	s2Arr := strings.Split(s2, ".")

	is.Equal(len(s1Arr), 3)
	is.Equal(len(s2Arr), 3)

}
