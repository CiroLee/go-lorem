package lorem

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiroLee/gear/gearmath"
	"github.com/CiroLee/gear/gearslice"
	"github.com/CiroLee/gear/gearstring"
)

// return a random hex color. set alpha true,will return a hex with alpha channel
func Hex(alpha bool) string {
	var hex string = "#"
	var length int = 6
	if alpha {
		length = 8
	}

	for i := 0; i < length; i++ {
		f, _ := randomFloat(0, 1, 6)
		hex += strconv.FormatInt(int64(f*16), 16)
	}
	return hex
}

// return a random rgb color array
func RgbArray() [3]int {
	return generateColorArray([2]int{0, 255}, [2]int{0, 255}, [2]int{0, 255})
}

// return a random rgba color array
func RgbaArray() [4]float32 {
	tmp := RgbArray()
	var rgb []int = make([]int, len(tmp))
	copy(rgb, tmp[:])
	rgb32 := gearslice.Map(rgb, func(el int, _ int) float32 {
		return float32(el)
	})
	alpha, _ := Float32(0, 1, 2)

	return [4]float32{rgb32[0], rgb32[1], rgb32[2], alpha}
}

// return a random rgb color string. support legacy and modern style
func RgbStr(style string) string {
	rgb := RgbArray()
	if style == "legacy" {
		return fmt.Sprintf("rgb(%d, %d, %d)", rgb[0], rgb[1], rgb[2])
	}
	return fmt.Sprintf("rgb(%d %d %d)", rgb[0], rgb[1], rgb[2])
}

// return a random rgba color string. support legacy and modern style
func RgbaStr(style string) string {
	rgba := RgbaArray()
	if style == "legacy" {
		return fmt.Sprintf("rgba(%v, %v, %v, %v)", rgba[0], rgba[1], rgba[2], rgba[3])
	} else {
		percent := math.Floor(float64(rgba[3] * 100))
		percentStr := strconv.FormatFloat(percent, 'f', -1, 64) + "%"
		return fmt.Sprintf("rgb(%v %v %v / %v)", rgba[0], rgba[1], rgba[2], percentStr)
	}
}

// return a random hsl color array. saturation and lightness will be decimal between 0 and 1(present percent)
func HslArray() [3]float32 {
	hue, _ := Float32(0, 360, 0)
	saturation, _ := Float32(0, 1, 2)
	lightness, _ := Float32(0, 1, 2)

	return [3]float32{hue, saturation, lightness}
}

// return a random hsla color array. saturation and lightness will be decimal between 0 and 1(present percent)
func HslaArray() [4]float32 {
	hsl := HslArray()
	alpha, _ := Float32(0, 1, 2)
	return [4]float32{hsl[0], hsl[1], hsl[2], alpha}
}

// return a random hsl color string. support legacy and modern style
func HslStr(style string) string {
	hsl := HslArray()
	hueStr := strconv.FormatFloat(float64(hsl[0]), 'f', -1, 64) + "deg"
	sat := math.Floor(float64(hsl[1] * 100))
	satStr := strconv.FormatFloat(sat, 'f', -1, 64) + "%"
	light := math.Floor(float64(hsl[2] * 100))
	lightStr := strconv.FormatFloat(light, 'f', -1, 64) + "%"
	if style == "legacy" {
		return fmt.Sprintf("hsl(%v, %v, %v)", hueStr, satStr, lightStr)
	}
	return fmt.Sprintf("hsl(%v %v %v)", hsl[0], satStr, lightStr)
}

// return a random hsla color string. support legacy and modern style
func HslaStr(style string) string {
	hsla := HslaArray()
	hueStr := strconv.FormatFloat(float64(hsla[0]), 'f', -1, 64) + "deg"
	sat := math.Floor(float64(hsla[1] * 100))
	satStr := strconv.FormatFloat(sat, 'f', -1, 64) + "%"
	light := math.Floor(float64(hsla[2] * 100))
	lightStr := strconv.FormatFloat(light, 'f', -1, 64) + "%"
	alpha := math.Floor(float64(hsla[3] * 100))
	alphaStr := strconv.FormatFloat(alpha, 'f', -1, 64) + "%"
	if style == "legacy" {
		return fmt.Sprintf("hsla(%v, %v, %v, %v)", hueStr, satStr, lightStr, alphaStr)
	}
	return fmt.Sprintf("hsl(%v %v %v / %v)", hueStr, satStr, lightStr, alphaStr)

}

func generateColorArray(red, green, blue [2]int) [3]int {
	var arr [3]int
	arr[0], _ = randomInteger(red[0], red[1])
	arr[1], _ = randomInteger(green[0], green[1])
	arr[2], _ = randomInteger(blue[0], blue[1])

	return arr
}

func extractRgb(rgb string) ([]float32, error) {
	var color []string
	rgbPrefixRe := regexp.MustCompile(`^(rgb|rgba|RGB|RGBA)\(`)
	legacyMode := regexp.MustCompile(`^(?:rgb|rgba|RGB|RGBA)\((\d{1,3}),(\d{1,3}),(\d{1,3})(?:, ?([\d.]+))?\)$`) // rgb(123,123,123)
	modernMode := regexp.MustCompile(`^(?:rgb|RGB)\((\s*\d{1,3}\s+){2}\s*\d{1,3}(?:\s*\/\s*([\d.]+%?))?\)$`)     // rgb(123 123 123)

	if !legacyMode.MatchString(rgb) && !modernMode.MatchString(rgb) {
		return []float32{0, 0, 0}, fmt.Errorf("rgb is an invalid rgb color")
	}
	tmp := rgbPrefixRe.ReplaceAllString(rgb, "")
	if legacyMode.MatchString(rgb) {
		color = strings.Split(strings.Trim(tmp, ")"), ",")
	} else {
		color = strings.Split(strings.Trim(tmp, ")"), " ")
		color = gearslice.Filter(color, func(el string, _ int) bool {
			return el != "/" && el != ""
		})
	}
	s := gearslice.Map(color, func(el string, _ int) float32 {
		if strings.Contains(el, "%") {
			float, _ := strconv.ParseFloat(strings.Replace(el, "%", "", -1), 32)
			return float32(float / 100)
		}
		float, _ := strconv.ParseFloat(el, 32)
		return float32(float)
	})

	return s, nil
}

func extractHSL(hsl string) ([]float32, error) {
	var color []string
	hslPrefixRe := regexp.MustCompile(`^(hsl|hsla|HSL|HSLA)\(`)
	legacyMode := regexp.MustCompile(`^(hsl|hsla|HSL|HSLA)\(((0|[1-9]\d*)deg),((100|[1-9]\d?)(\.\d+)?%),((100|[1-9]\d?)(\.\d+)?%)(,(0\.\d+|1|0))? *\)$`)
	modernMode := regexp.MustCompile(`^(hsl|HSL)\((0?\d{1,2}|1[0-9]{2}|2[0-9]{1,2}|3[0-5]\d|360)deg\s(0?\d{1,2}|100)%\s(0?\d{1,2}|100)%(\s/\s(0?\d{1,2}|100)%)?\)$`)

	if !legacyMode.MatchString(hsl) && !modernMode.MatchString(hsl) {
		return []float32{0, 0, 0}, fmt.Errorf("hsl is an invalid hsl color")
	}
	tmp := hslPrefixRe.ReplaceAllString(hsl, "")
	if legacyMode.MatchString(hsl) {
		color = strings.Split(strings.Trim(tmp, ")"), ",")
	} else {
		color = strings.Split(strings.Trim(tmp, ")"), " ")
		color = gearslice.Filter(color, func(el string, _ int) bool {
			return el != "/" && el != ""
		})
	}
	s := gearslice.Map(color, func(el string, index int) float32 {
		el = strings.Replace(el, "deg", "", -1)
		if strings.Contains(el, "%") && index == 3 {
			float, _ := strconv.ParseFloat(strings.Replace(el, "%", "", -1), 32)
			return float32(float / 100)
		}
		float, _ := strconv.ParseFloat(strings.Replace(el, "%", "", -1), 32)
		return float32(float)
	})
	return s, nil
}

func hexToRgb(hex string) ([3]float32, error) {
	re := regexp.MustCompile("^#([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$")
	hex = strings.ToLower(hex)
	if !re.MatchString(hex) {
		return [3]float32{0, 0, 0}, fmt.Errorf("hex is an invalid hex color")
	}
	if len(hex) == 4 {
		_hex := gearstring.Substring(hex, 1, len(hex))
		hex = "#" + strings.Repeat(_hex, 2)
	}
	var rgb []float32
	for i := 1; i < 7; i += 2 {
		str16 := gearstring.Substring(hex, i, i+2)
		c, _ := strconv.ParseInt(str16, 16, 32)
		fmt.Printf("c: %v\n", c)
		rgb = append(rgb, float32(c))
	}
	return [3]float32{rgb[0], rgb[1], rgb[2]}, nil
}

func rgbToHex(rgb [3]float32) string {

	r := int(rgb[0]) << 16
	g := int(rgb[1]) << 8
	b := int(rgb[2])
	return "#" + strconv.FormatInt(int64(r+g+b), 16)
}

func rgbToHsl(rgb [3]float32) [3]float32 {
	r := float64(rgb[0] / 255)
	g := float64(rgb[1] / 255)
	b := float64(rgb[2] / 255)
	max := gearmath.Max([]float64{r, g, b})
	min := gearmath.Min([]float64{r, g, b})
	mid := (max + min) / 2

	var s float64
	h, l := mid, mid
	if min == max {
		h = 0.0
	} else {
		var delta = max - min
		if l > 0.5 {
			s = delta / (2.0 - max - min)
		} else {
			s = delta / (min + max)
		}
		switch max {
		case r:
			var temp float64
			if g < b {
				temp = 6.0
			} else {
				temp = 0.0
			}
			h = (g-b)/delta + temp
		case g:
			h = (b-r)/delta + 2
		case b:
			h = (r-g)/delta + 4
		}
	}

	h *= 60

	return [3]float32{
		float32(math.Round(h)),
		float32(math.Round(s*100) / 100),
		float32(math.Round(l*100) / 100),
	}
}
