package lorem

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CiroLee/gear/gearstring"
)

const placeholder_image_url = "https://dummyimage.com"
const picsum_image_url = "https://picsum.photos"
const loremflicker_url = "https://loremflickr.com"

type Hsl = [3]float32
type Size struct {
	Width  int
	Height int
}
type PlaceholderOption struct {
	Width     int
	Height    int
	Text      string
	BgColor   string
	FontColor string
}

func isDark(hsl Hsl) bool {
	return hsl[2] < 0.5
}

func initSize(option Size) Size {
	if option.Width > 0 && option.Height <= 0 {
		option.Height = option.Width
	} else if option.Width <= 0 && option.Height > 0 {
		option.Width = option.Height
	} else if option.Width <= 0 && option.Height <= 0 {
		option.Width, _ = randomInteger(320, 1024)
		option.Height, _ = randomInteger(320, 1024)
	}
	return Size{
		Width:  option.Width,
		Height: option.Height,
	}
}

func hexToHsl(hex string) ([3]float32, error) {
	rgb, err := hexToRgb(hex)
	if err != nil {
		return rgb, err
	}
	return rgbToHsl(rgb), nil
}

// return a random placeholder image with pure color background
func Placeholder(option PlaceholderOption) (string, error) {
	if option.BgColor != "" && option.FontColor != "" && (!isHex(option.BgColor) || !isHex(option.FontColor)) {
		return "", fmt.Errorf("BgColor or FontColor is not hex format color")
	}
	if option.BgColor == "" {
		option.BgColor = Hex(false)
	}
	var size Size = initSize(Size{Width: option.Width, Height: option.Height})
	bgHsl, _ := hexToHsl(option.BgColor)
	dark := isDark(bgHsl)
	if option.FontColor == "" {
		if dark {
			option.FontColor = "#ffffff"
		} else {
			option.FontColor = "#000000"
		}
	}
	imgUrl := gearstring.Contact(placeholder_image_url, "/", strconv.Itoa(size.Width), "x", strconv.Itoa(size.Height), "/", strings.TrimPrefix(option.BgColor, "#"))

	if option.Text != "" {
		return gearstring.Contact(imgUrl, "/", strings.TrimPrefix(option.FontColor, "#"), "&text=", option.Text), nil
	}
	return imgUrl, nil
}
