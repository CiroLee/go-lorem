package lorem

import (
	"strconv"
	"strings"

	"github.com/CiroLee/go-lorem/data"
)

// return a random ipv4 address
func Ipv4() string {
	ip := ""
	for i := 0; i < 4; i++ {
		n, _ := randomInteger(0, 255)
		ip += strconv.Itoa(n) + "."
	}
	return strings.TrimRight(ip, ".")
}

// return a random ipv6 address
func Ipv6() string {
	var r []string
	hash := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	for i := 0; i < 8; i++ {
		element := Elements(hash, 4)
		r = append(r, strings.Join(element, ""))
	}
	return strings.Join(r, ":")
}

// return a random top-level-domain
func Tld() (tld, tldType string) {
	tldType = randomElement([]string{"iTLD", "ccTLD"})
	tldData := data.DOMAINS[tldType]
	index, _ := randomInteger(0, len(tldData)-1)
	tld = tldData[index]
	return
}
