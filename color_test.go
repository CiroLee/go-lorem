package lorem

import (
	"strings"
	"testing"

	"github.com/CiroLee/gear/gearslice"
	"github.com/stretchr/testify/assert"
)

func TestGenerateColorArray(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	color := generateColorArray([2]int{0, 255}, [2]int{0, 255}, [2]int{0, 255})
	colorSlice := make([]int, len(color))
	copy(colorSlice, color[:])
	is.Len(color, 3)
	is.True(gearslice.Every(colorSlice, func(el int, _ int) bool {
		return el >= 0 && el <= 255
	}))
}

func TestHex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	hex1 := Hex(false)
	hex2 := Hex(true)

	is.Len(hex1, 7)
	is.Len(hex2, 9)
	is.True(strings.HasPrefix(hex1, "#"))
	is.True(strings.HasPrefix(hex2, "#"))
}
func TestRGBArray(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	color := RgbArray()
	colorSlice := make([]int, len(color))
	copy(colorSlice, color[:])
	is.Len(color, 3)
	is.True(gearslice.Every(colorSlice, func(el int, _ int) bool {
		return el >= 0 && el <= 255
	}))
}

func TestRGBAArray(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	color := RgbaArray()
	colorSlice := make([]float32, len(color))
	copy(colorSlice, color[:])
	is.Len(color, 4)
	is.True(gearslice.Every(colorSlice, func(el float32, i int) bool {
		if i < 4 {
			return el >= 0 && el <= 255
		} else {
			return el < 1
		}
	}))
}

func TestRGBStr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rgbstr1 := RgbStr("legacy")
	rgbstr2 := RgbStr("modern")

	is.Equal(strings.Count(rgbstr1, ","), 2)
	is.True(strings.HasPrefix(rgbstr1, "rgb"))
	is.Equal(strings.Count(rgbstr2, " "), 2)
	is.True(strings.HasPrefix(rgbstr2, "rgb"))
}

func TestRGBAStr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rgbastr1 := RgbaStr("legacy")
	rgbastr2 := RgbaStr("modern")
	is.Equal(strings.Count(rgbastr1, ","), 3)
	is.True(strings.HasPrefix(rgbastr1, "rgba"))
	is.Equal(strings.Count(rgbastr2, "/"), 1)
	is.True(strings.HasPrefix(rgbastr2, "rgb"))
}

func TestHSLArray(t *testing.T) {
	is := assert.New(t)

	hsl := HslArray()
	is.LessOrEqual(hsl[0], float32(360))
	is.GreaterOrEqual(hsl[0], float32(0))
	is.LessOrEqual(hsl[1], float32(1))
	is.GreaterOrEqual(hsl[1], float32(0))
	is.LessOrEqual(hsl[2], float32(1))
	is.GreaterOrEqual(hsl[2], float32(0))
}

func TestHSLAArray(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	hsla := HslaArray()
	is.LessOrEqual(hsla[0], float32(360))
	is.GreaterOrEqual(hsla[0], float32(0))
	is.LessOrEqual(hsla[1], float32(1))
	is.GreaterOrEqual(hsla[1], float32(0))
	is.LessOrEqual(hsla[2], float32(1))
	is.GreaterOrEqual(hsla[2], float32(0))
	is.LessOrEqual(hsla[3], float32(1))
	is.GreaterOrEqual(hsla[3], float32(0))
}

func TestHSLStr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	hsl1 := HslStr("legacy")
	hsl2 := HslStr("modern")

	is.True(strings.HasPrefix(hsl1, "hsl"))
	is.Equal(strings.Count(hsl1, ","), 2)
	is.True(strings.HasPrefix(hsl2, "hsl"))
	is.Equal(strings.Count(hsl2, " "), 2)
}

func TestHSLAStr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	hsla1 := HslaStr("legacy")
	hsla2 := HslaStr("modern")

	is.True(strings.HasPrefix(hsla1, "hsla"))
	is.Equal(strings.Count(hsla1, ","), 3)
	is.True(strings.HasPrefix(hsla2, "hsl"))
	is.Equal(strings.Count(hsla2, "/"), 1)
}
