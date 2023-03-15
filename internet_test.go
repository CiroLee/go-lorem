package lorem

import (
	"strings"
	"testing"

	"github.com/CiroLee/gear/gearslice"
	"github.com/stretchr/testify/assert"
)

func TestIpv4(t *testing.T) {
	is := assert.New(t)

	ipv4 := Ipv4()
	l := strings.Split(ipv4, ".")
	is.Len(l, 4)
}

func TestIpv6(t *testing.T) {
	is := assert.New(t)

	ipv6 := Ipv6()
	l := strings.Split(ipv6, ":")
	is.Len(l, 8)
}

func TestTld(t *testing.T) {
	is := assert.New(t)
	tld, tldType := Tld()

	is.True(gearslice.Includes([]string{"iTLD", "ccTLD"}, tldType))
	is.True(strings.HasPrefix(tld, "."))
}

func TestDomain(t *testing.T) {
	is := assert.New(t)

	d := Domain(2)
	is.Equal(strings.Count(d, "."), 2)
}

func TestSubDirectory(t *testing.T) {
	is := assert.New(t)
	sub := subDirectory(2)

	is.Equal(strings.Count(sub, "/"), 2)
}

func TestUrl(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	url1 := Url(UrlOption{
		Protocol: "https",
	})
	url2 := Url(UrlOption{
		Suffix: ".com",
	})
	url3 := SimpleUrl()

	is.True(strings.HasPrefix(url1, "https"))
	is.Contains(url2, ".com")
	is.Contains(url3, "https")
}

func TestEmail(t *testing.T) {
	is := assert.New(t)

	email := Email()
	is.Contains(email, "@")
}
