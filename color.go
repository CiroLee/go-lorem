package lorem

import (
	"fmt"
	"math"
	"strconv"

	"github.com/CiroLee/gear/gearslice"
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
	sat := math.Floor(float64(hsl[1] * 100))
	satStr := strconv.FormatFloat(sat, 'f', -1, 64) + "%"
	light := math.Floor(float64(hsl[2] * 100))
	lightStr := strconv.FormatFloat(light, 'f', -1, 64) + "%"
	if style == "legacy" {
		return fmt.Sprintf("hsl(%v, %v, %v)", hsl[0], satStr, lightStr)
	}
	return fmt.Sprintf("hsl(%v %v %v)", hsl[0], satStr, lightStr)
}

// return a random hsla color string. support legacy and modern style
func HslaStr(style string) string {
	hsla := HslaArray()
	sat := math.Floor(float64(hsla[1] * 100))
	satStr := strconv.FormatFloat(sat, 'f', -1, 64) + "%"
	light := math.Floor(float64(hsla[2] * 100))
	lightStr := strconv.FormatFloat(light, 'f', -1, 64) + "%"
	alpha := math.Floor(float64(hsla[3] * 100))
	alphaStr := strconv.FormatFloat(alpha, 'f', -1, 64) + "%"
	if style == "legacy" {
		return fmt.Sprintf("hsla(%v, %v, %v, %v)", hsla[0], satStr, lightStr, alphaStr)
	}
	return fmt.Sprintf("hsl(%v %v %v / %v)", hsla[0], satStr, lightStr, alphaStr)

}

func generateColorArray(red, green, blue [2]int) [3]int {
	var arr [3]int
	arr[0], _ = randomInteger(red[0], red[1])
	arr[1], _ = randomInteger(green[0], green[1])
	arr[2], _ = randomInteger(blue[0], blue[1])

	return arr
}
