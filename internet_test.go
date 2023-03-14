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
	is.Equal(len(l), 4)
}

func TestIpv6(t *testing.T) {
	is := assert.New(t)

	ipv6 := Ipv6()
	l := strings.Split(ipv6, ":")
	is.Equal(len(l), 8)
}

func TestTld(t *testing.T) {
	is := assert.New(t)
	tld, tldType := Tld()

	is.True(gearslice.Includes([]string{"iTLD", "ccTLD"}, tldType))
	is.True(strings.HasPrefix(tld, "."))
}
