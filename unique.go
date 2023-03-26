package lorem

import (
	"strings"

	"github.com/CiroLee/gear/gearslice"
	"github.com/CiroLee/go-lorem/data"
)

// return a random Chinese vehicle registration mark
func Vrm() string {
	tmp := gearslice.Filter(data.ALPHABET, func(el string, _ int) bool {
		return el != "i" && el != "0"
	})
	vrmAlphabet := strings.Join(tmp, "")
	vrmAlphabet = strings.ToUpper(vrmAlphabet)
	trailTag := "挂学警港澳"
	boolean := func() bool {
		return randomElement([]bool{true, false})
	}()
	var body string
	if boolean {
		body = StrBy(4, vrmAlphabet+data.NUMBER) + StrBy(1, trailTag)
	} else {
		body = StrBy(4, vrmAlphabet+data.NUMBER)
	}
	return StrBy(1, data.RegionAbbr) + StrBy(1, vrmAlphabet) + body
}
