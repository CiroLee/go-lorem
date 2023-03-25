package lorem

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDark(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := isDark([3]float32{124, 0.23, 0.6})
	r2 := isDark([3]float32{124, 0.23, 0.23})

	is.False(r1)
	is.True(r2)
}

func TestInitSize(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	size1 := initSize(Size{100, 100})
	size2 := initSize(Size{Width: 100})
	size3 := initSize(Size{Height: 100})
	size4 := initSize(Size{})

	is.Equal(size1, Size{100, 100})
	is.Equal(size2, Size{100, 100})
	is.Equal(size3, Size{100, 100})
	is.LessOrEqual(size4.Width, uint(1024))
	is.LessOrEqual(size4.Height, uint(1024))
	is.GreaterOrEqual(size4.Height, uint(320))
	is.GreaterOrEqual(size4.Width, uint(320))
}

func TestHexToHsl(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	hsl1, err1 := hexToHsl("#fff")
	hsl2, err2 := hexToHsl("#ff0000a8")

	is.Equal(hsl1, [3]float32{0, 0, 1})
	is.NoError(err1)
	is.Equal(hsl2, [3]float32{0, 0, 0})
	is.Error(err2)
}

func TestPlaceholder(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	p1, err1 := Placeholder(PlaceholderOption{BgColor: "1", FontColor: "2"})
	p2, err2 := Placeholder(PlaceholderOption{BgColor: "#fff", FontColor: "3"})
	p3, err3 := Placeholder(PlaceholderOption{BgColor: "3", FontColor: "#333"})
	p4, err4 := Placeholder(PlaceholderOption{FontColor: "#333"})
	p5, err5 := Placeholder(PlaceholderOption{BgColor: "#fff"})
	p6, err6 := Placeholder(PlaceholderOption{Text: "color"})

	is.Empty(p1)
	is.Error(err1)
	is.Empty(p2)
	is.Error(err2)
	is.Empty(p3)
	is.Error(err3)
	is.NotEmpty(p4)
	is.NoError(err4)
	is.Contains(p5, "fff")
	is.NoError(err5)
	is.Contains(p6, "text=color")
	is.NoError(err6)
}

func TestSimplePlaceholder(t *testing.T) {
	is := assert.New(t)
	p := SimplePlaceholder()

	is.Contains(p, "text=image")
}

func TestPicsum(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	p1 := Picsum(PicsumOption{})
	p2 := Picsum(PicsumOption{Width: 200})
	p3 := Picsum(PicsumOption{Grayscale: true})
	p4 := Picsum(PicsumOption{Blur: 2, Grayscale: true})
	p5 := Picsum(PicsumOption{Blur: 2, Grayscale: false})
	p6 := SimplePicsum()

	is.Equal(strings.Count(p1, "/"), 4)
	is.Equal(strings.Count(p2, "200"), 2)
	is.True(strings.HasSuffix(p3, "?grayscale"))
	is.Contains(p4, "?grayscale")
	is.True(strings.HasSuffix(p4, "&blur=2"))
	is.True(strings.HasSuffix(p5, "?blur=2"))
	is.Equal(strings.Count(p6, "/"), 4)
}
