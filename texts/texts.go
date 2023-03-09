package lorem

import (
	"github.com/CiroLee/go-lorem/constants"
	"github.com/CiroLee/go-lorem/core"
	"github.com/CiroLee/go-lorem/utils"
)

type WordProps struct {
	Num   uint
	Lang  string
	Upper bool
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

func Word(props WordProps) string {
	var str string
	var length uint = 2
	if props.Num > 0 {
		length = props.Num
	}
	for i := 0; i < int(length); i++ {
		str += Letter(props.Lang)
	}

	if props.Upper {
		return utils.Capitalize(str)
	}
	return str
}
