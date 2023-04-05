package lorem

import (
	"regexp"
	"strings"
	"testing"

	"github.com/CiroLee/gear/gearslice"
	"github.com/CiroLee/go-lorem/data"
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

func TestMobile(t *testing.T) {
	is := assert.New(t)

	m := Mobile()
	is.Len(m, 11)
}

func TestMobileHideMiddle(t *testing.T) {
	is := assert.New(t)

	m := MobileHideMiddle()
	is.Len(m, 11)
	is.Equal(strings.Count(m, "*"), 4)
}

func TestLandline(t *testing.T) {
	is := assert.New(t)

	n := Landline()
	r := strings.Split(n, "-")[1]

	is.LessOrEqual(len(r), 8)
	is.GreaterOrEqual(len(r), 7)
}

func TestHttpMethod(t *testing.T) {
	is := assert.New(t)

	m := HttpMethod()
	is.True(gearslice.Includes(data.HTTP_METHODS, m))
}

func TestHttpStatusCode(t *testing.T) {
	is := assert.New(t)

	code := HttpStatusCode()
	is.LessOrEqual(code, 599)
}

func TestUUID(t *testing.T) {
	is := assert.New(t)

	uuid := UUID()
	match, err := regexp.MatchString("[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}", uuid)
	is.True(match)
	is.Nil(err)

	uuids := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		uuid := UUID()
		assert.False(t, uuids[uuid], "duplicate UUID generated: %s", uuid)
		uuids[uuid] = true
	}
}

func TestProtocol(t *testing.T) {
	is := assert.New(t)
	p := Protocol()

	is.True(gearslice.Includes(data.PROTOCOL, p))
}
