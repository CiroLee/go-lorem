package lorem

import (
	"strings"
	"testing"

	"github.com/CiroLee/gear/gearslice"
	"github.com/stretchr/testify/assert"
)

func TestAmount(t *testing.T) {
	is := assert.New(t)

	a := Amount(10000, 999999, 4, "$")
	b := Amount(99999, 1000, 4, "")
	c := Amount(10000, 999999, 0, "$")

	tmp := strings.TrimLeft(a, "$")
	intPart := strings.Split(tmp, ".")[0]
	decimalPart := strings.Split(tmp, ".")[1]

	is.True(strings.HasPrefix(a, "$"))
	is.LessOrEqual(len(intPart), 7)
	is.GreaterOrEqual(len(intPart), 5)
	is.Len(decimalPart, 4)
	is.Empty(b)
	is.False(strings.Contains(c, "."))
}

func TestGetIssuer(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	v1 := getIssuer("")
	v2 := getIssuer("abc")
	v3 := getIssuer(UnionPay)
	v4 := getIssuer(Amex)

	is.Empty(v1)
	is.Empty(v2)
	is.Equal(v3, "62")
	is.True(gearslice.Includes([]string{"34", "37"}, v4))
}

func TestBankCardNumber(t *testing.T) {
	is := assert.New(t)

	num := BankCardNumber(UnionPay)

	is.True(strings.HasPrefix(num, "62"))
}
