package lorem

import (
	"strconv"
	"strings"

	"github.com/CiroLee/go-lorem/data"
)

var issuerMap = map[string][]int{
	"visa":       {4},
	"amex":       {34, 37},
	"jcb":        {35},
	"mastercard": {51, 52, 53, 54, 55},
	"unionpay":   {62},
}

const (
	Visa       = "visa"
	Amex       = "amex"
	Jcb        = "jcb"
	MasterCard = "mastercard"
	UnionPay   = "unionpay"
)

// return a random amount with special prefix
func Amount(min, max float64, decimal uint, prefix string) string {
	num, err := randomFloat(min, max, decimal)
	if err != nil {
		return ""
	}
	str := strconv.FormatFloat(num, 'f', -1, 64)
	dotIndex := strings.Index(str, ".")
	if dotIndex == -1 {
		dotIndex = len(str)
	}
	intPart := str[:dotIndex]
	var decimalPart string
	if dotIndex < len(str)-1 {
		decimalPart = str[dotIndex+1:]
	}
	if len(decimalPart) < int(decimal) {
		decimalPart += strings.Repeat("0", int(decimal)-len(decimalPart))
	}

	// format intPart
	var formattedInt string
	for i, v := range intPart {
		if i > 0 && (len(intPart)-i)%3 == 0 {
			formattedInt += ","
		}
		formattedInt += string(v)
	}
	// combine intPart and decimalPart
	formattedNum := formattedInt
	if decimalPart != "" {
		formattedNum += "." + decimalPart
	}
	return prefix + formattedNum
}

// return a random bank card number. support Visa, Amex, Jcb, MasterCard and UnionPay card issuer
func BankCardNumber(issuer string) string {
	issuerNum := getIssuer(issuer)
	min := 13 - len(issuerNum)
	max := 19 - len(issuerNum)
	random, _ := randomInteger(min, max)
	var rest = StrBy(uint(random), data.NUMBER_STR)
	return issuerNum + rest
}
func getIssuer(issuer string) string {
	s, ok := issuerMap[issuer]
	if !ok {
		return ""
	}
	var val int
	if len(s) == 1 {
		val = s[0]
	} else {
		val = randomElement(s)
	}
	return strconv.Itoa(val)
}
