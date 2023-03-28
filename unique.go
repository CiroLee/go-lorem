package lorem

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/CiroLee/gear/gearslice"
	"github.com/CiroLee/gear/gearstring"
	"github.com/CiroLee/go-lorem/consts"
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
		body = StrBy(4, vrmAlphabet+data.NUMBER_STR) + StrBy(1, trailTag)
	} else {
		body = StrBy(4, vrmAlphabet+data.NUMBER_STR)
	}
	return StrBy(1, data.RegionAbbr) + StrBy(1, vrmAlphabet) + body
}

// return a random vehicle identification number
func Vin() string {
	re := regexp.MustCompile(`i|o|q`)
	vinAlphabet := re.ReplaceAllString(strings.Join(data.ALPHABET, ""), "")
	vinAlphabet = strings.ToUpper(vinAlphabet)
	randomInt, _ := IntBy(2, true)
	wmi := StrBy(3, "123469"+vinAlphabet)
	vds := StrBy(2, vinAlphabet) + strconv.Itoa(randomInt) + StrBy(1, vinAlphabet) + StrBy(1, data.NUMBER_STR+"X")
	vis := StrBy(1, regexp.MustCompile(`u|z`).ReplaceAllString(vinAlphabet, "")+gearstring.Substring(data.NUMBER_STR, 1, len(data.NUMBER_STR))) + StrBy(1, vinAlphabet) + StrBy(6, data.NUMBER_STR)

	return wmi + vds + vis
}

// return a random Chinese Mainland ID
func Id() string {
	pcNum, _ := IntBy(6, true)
	start := time.Date(1970, time.Month(1), 1, 0, 0, 0, 0, time.Now().Location())
	end := time.Date(2999, time.Month(12), 31, 0, 0, 0, 0, time.Now().Location())
	date := FormatDate(start, end, consts.NoSeparatorLayout)
	boolean := func() bool {
		return randomElement([]bool{true, false})
	}()

	var trailNum string
	if boolean {
		trailNum = StrBy(4, data.NUMBER_STR)
	} else {
		trailNum = StrBy(3, data.NUMBER_STR) + "x"
	}

	return strconv.Itoa(pcNum) + date + trailNum
}
