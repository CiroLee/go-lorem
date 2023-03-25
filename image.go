package lorem

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiroLee/gear/gearstring"
)

const placeholder_image_url = "https://dummyimage.com"
const picsum_image_url = "https://picsum.photos"
const loremflicker_url = "https://loremflickr.com"

type Hsl = [3]float32
type Size struct {
	Width  uint
	Height uint
}
type PlaceholderOption struct {
	Width     uint
	Height    uint
	Text      string
	BgColor   string
	FontColor string
}

type PicsumOption struct {
	Width     uint
	Height    uint
	Grayscale bool
	Blur      uint
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
		w, _ := randomInteger(320, 1024)
		h, _ := randomInteger(320, 1024)
		option.Width = uint(w)
		option.Height = uint(h)
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

// return a random placeholder image with pure color background with special struct
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
	imgUrl := gearstring.Contact(placeholder_image_url, "/", strconv.Itoa(int(size.Width)), "x", strconv.Itoa(int(size.Height)), "/", strings.TrimPrefix(option.BgColor, "#"))

	if option.Text != "" {
		return gearstring.Contact(imgUrl, "/", strings.TrimPrefix(option.FontColor, "#"), "&text=", option.Text), nil
	}
	return imgUrl, nil
}

// return a random placeholder image without params
func SimplePlaceholder() string {
	img, _ := Placeholder(PlaceholderOption{Text: "image"})
	return img
}

// return a random image from unSplash by special param struct
func Picsum(option PicsumOption) string {
	var size = initSize(Size{option.Width, option.Height})
	re := regexp.MustCompile(`\/?grayscale$`)
	var blur string
	if option.Blur >= 10 {
		blur = "10"
	} else {
		blur = strconv.Itoa(int(option.Blur))
	}
	imgUrl := gearstring.Contact(picsum_image_url, "/", strconv.Itoa(int(size.Width)), "/", strconv.Itoa(int(size.Height)))
	if option.Grayscale {
		imgUrl += "?grayscale"
	}

	if option.Blur > 0 {
		if re.MatchString(imgUrl) {
			imgUrl += "&blur=" + blur
		} else {
			imgUrl += "?blur=" + blur
		}
	}
	return imgUrl
}

// return a random image. simple use of Picsum without params
func SimplePicsum() string {
	return Picsum(PicsumOption{})
}
