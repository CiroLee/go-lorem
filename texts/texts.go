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
