package lorem

import (
	"strings"

	"github.com/CiroLee/go-lorem/constants"
	"github.com/CiroLee/go-lorem/core"
	"github.com/CiroLee/go-lorem/utils"
)

const BASE_NUM = 14

type BaseProps struct {
	Num  uint
	Lang string
}

// return a random letter
func Letter(lang string) string {
	var source []string
	if lang == "en" {
		source = constants.ALPHABET
	} else {
		source = constants.ZH_CHARACTERS
	}
	index, _ := core.RandomInteger(0, len(source)-1)
	return source[index]
}

// return a random word
func Word(props BaseProps) string {
	var str string
	var length uint = 2
	if props.Num > 0 {
		length = props.Num
	}
	for i := 0; i < int(length); i++ {
		str += Letter(props.Lang)
	}
	return str
}

// return a random sentence
func Sentence(props BaseProps) string {
	var str string
	var length = 1
	trail := "，"
	if props.Num > 0 {
		length = int(props.Num)
	}
	if props.Lang == "en" {
		trail = " "
	}
	for i := 0; i < length; i++ {
		num, _ := core.RandomInteger(2, BASE_NUM)
		str += Word(BaseProps{
			Lang: props.Lang,
			Num:  uint(num),
		}) + trail
	}
	return utils.Capitalize(strings.TrimRight(str, trail))
}

// return a random paragraph
func Paragraph(props BaseProps) string {
	var str string
	var length = 1
	trail := "。"
	if props.Num > 0 {
		length = int(props.Num)
	}
	if props.Lang == "en" {
		trail = "."
	}
	for i := 0; i < length; i++ {
		num, _ := core.RandomInteger(1, BASE_NUM)
		str += Sentence(BaseProps{
			Num:  uint(num),
			Lang: props.Lang,
		}) + trail
	}
	return str
}

// return a name
func Name(lang string, upper bool) string {
	var source []string
	if lang == "en" {
		source = constants.COMMON_EN_NAMES
	} else {
		source = constants.COMMON_ZH_NAMES
	}
	index, _ := core.RandomInteger(0, len(source)-1)
	if upper {
		return utils.Capitalize(source[index])
	}
	return source[index]
}

func Str(length uint) string {
	var str string
	s := []rune(constants.STRINGS)
	lens := 4
	if length > 0 {
		lens = int(length)
	}
	for i := 0; i < lens; i++ {
		index, _ := core.RandomInteger(0, len(s)-1)
		str += string(s[index])
	}
	return str
}
