package lorem

import (
	"strings"

	"github.com/CiroLee/gear/gearstring"
	"github.com/CiroLee/go-lorem/data"
)

const baseNum = 10

// return a random letter
func Letter(lang string) string {
	var source []string
	if lang == "en" {
		source = data.ALPHABET
	} else {
		source = data.ZH_CHARACTERS
	}
	index, _ := randomInteger(0, len(source)-1)
	return source[index]
}

// return a random word
func Word(num uint, lang string) string {
	var str string
	var length uint = 2
	if num > 0 {
		length = num
	}
	for i := 0; i < int(length); i++ {
		str += Letter(lang)
	}
	return str
}

// return a random sentence
func Sentence(num uint, lang string) string {
	var str string
	var length uint = 1
	trail := "，"
	if num > 0 {
		length = num
	}
	if lang == "en" {
		trail = " "
	}
	for i := 0; i < int(length); i++ {
		num, _ := randomInteger(2, baseNum)
		str += Word(uint(num), lang) + trail
	}
	return gearstring.ToUpperCase(strings.TrimRight(str, trail))
}

// return a random paragraph
func Paragraph(num uint, lang string) string {
	var str string
	var length uint = 1
	trail := "。"
	if num > 0 {
		length = num
	}
	if lang == "en" {
		trail = "."
	}
	for i := 0; i < int(length); i++ {
		num, _ := randomInteger(1, baseNum)
		str += Sentence(uint(num), lang) + trail
	}
	return str
}

// return a name
func Name(lang string, upper bool) string {
	var source []string
	if lang == "en" {
		source = data.COMMON_EN_NAMES
	} else {
		source = data.COMMON_ZH_NAMES
	}
	index, _ := randomInteger(0, len(source)-1)
	if upper {
		return gearstring.ToUpperCase(source[index])
	}
	return source[index]
}

// return a random string
func Str(length uint) string {
	var str string
	s := []rune(data.STRINGS)
	lens := 4
	if length > 0 {
		lens = int(length)
	}
	for i := 0; i < lens; i++ {
		index, _ := randomInteger(0, len(s)-1)
		str += string(s[index])
	}
	return str
}
