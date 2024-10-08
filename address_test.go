package lorem

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomCity(t *testing.T) {
	is := assert.New(t)

	city := randomCity("120000")

	is.Equal(city.Name, "天津市")
	is.Equal(city.Code, "120000")
}

func TestRandomCounty(t *testing.T) {
	is := assert.New(t)
	county := randomCounty("110000", "110000")
	is.NotEmpty(county.Code)
	is.NotEmpty(county.Name)
}

func TestProvince(t *testing.T) {
	is := assert.New(t)
	p := Province()
	is.NotEmpty(p.Code)
	is.NotEmpty(p.Name)
}

func TestCity(t *testing.T) {
	is := assert.New(t)
	c := City()
	is.NotEmpty(c.Code)
	is.NotEmpty(c.Code)
}

func TestCounty(t *testing.T) {
	is := assert.New(t)
	c := County()

	is.Equal(reflect.TypeOf(c.Code).Kind(), reflect.String)
	is.Equal(reflect.TypeOf(c.Name).Kind(), reflect.String)
}

func TestRoad(t *testing.T) {
	is := assert.New(t)
	r := Road()
	is.NotEmpty(r)
}

func TestAddress(t *testing.T) {
	is := assert.New(t)
	addressGap := Address(true)
	addressNoGap := Address(false)

	length := len(strings.Split(addressGap, " "))
	is.LessOrEqual(length, 3)
	is.NotEmpty(addressNoGap)
}

func TestFullAddress(t *testing.T) {
	is := assert.New(t)
	addressGap := FullAddress(true)
	addressNoGap := FullAddress(false)

	length := len(strings.Split(addressGap, " "))
	is.LessOrEqual(length, 5)
	is.NotEmpty(addressNoGap)
}

func TestZipCode(t *testing.T) {
	is := assert.New(t)
	zip := strconv.Itoa(ZipCode())
	is.Len(zip, 6)
}

func TestDegToDms(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	v1 := degToDms(192.8765)
	v2 := degToDms(-192.8765)

	is.Equal(v1, "192°52′35″")
	is.Equal(v2, "-192°52′35″")
}

func TestLongAndLat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	v1 := LongAndLat("deg")
	v2 := LongAndLat("dms")

	for _, v := range v1 {
		is.Contains(v, "°")
	}
	for _, v := range v2 {
		is.Contains(v, "°")
		is.Contains(v, "′")
		is.True(strings.HasSuffix(v, "″"))
	}
}
