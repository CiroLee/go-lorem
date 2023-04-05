package lorem

import (
	cryptoRand "crypto/rand"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiroLee/gear/gearstring"
	"github.com/CiroLee/go-lorem/data"
)

type UrlOption struct {
	SubDir      uint
	Protocol    string
	Suffix      string
	DomainLevel uint
}

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

// return a random domain at special level
func Domain(level uint) string {
	var str string
	source := "0123456789abcdefghijklmnopqrstuvwxyz-"
	for i := 0; i < int(level); i++ {
		random, _ := randomInteger(1, 10)
		re := regexp.MustCompile("^-|-$")
		tmp := StrBy(uint(random), source)
		tmp = string(re.ReplaceAllString(tmp, Letter("en"))) + "."
		str += tmp
	}
	tld, _ := Tld()
	return strings.TrimRight(str, ".") + tld
}

// return a random url via the option struct
/*
 type UrlOption struct {
	SubDir       uint      layers of sub-directory,default is between [1,4]
	Protocol     string    protocol of url,default is random
	Suffix       string    suffix of the url, such as .com, .org,default is random
	DomainLevel  uint      level of sub domain, default is between [1,3]
 }
*/
func Url(option UrlOption) string {
	if option.SubDir == 0 {
		subDir, _ := randomInteger(1, 4)
		option.SubDir = uint(subDir)
	}
	if option.DomainLevel == 0 {
		random, _ := randomInteger(1, 3)
		option.DomainLevel = uint(random)
	}
	if option.Suffix == "" {
		option.Suffix, _ = Tld()
	}
	if option.Protocol == "" {
		option.Protocol = randomElement(data.PROTOCOL)
	}
	subDir := subDirectory(option.SubDir)
	prefix := Domain(option.DomainLevel)
	return option.Protocol + "://" + gearstring.Substring(prefix, 0, strings.LastIndex(prefix, ".")) + option.Suffix + subDir
}

// return a random url without params. if you want to set please use Url
func SimpleUrl() string {
	return Url(UrlOption{
		Protocol: "https",
		Suffix:   ".com",
	})
}

// return a random web protocol
func Protocol() string {
	return randomElement(data.PROTOCOL)
}

// return a random email
func Email() string {
	random, _ := randomInteger(1, data.BaseNum)
	name := StrBy(uint(random), data.NUMBER_ALPHABET)
	emailSuffix := randomElement(data.EMAIL_SUFFIX)
	return name + "/" + emailSuffix
}

// return 11-digit Mainland China mobile number
func Mobile() string {
	start := randomElement(data.MESH_NUM)
	var rest string
	for i := 0; i < 8; i++ {
		random, _ := randomInteger(0, 9)
		rest += strconv.Itoa(random)
	}
	return strconv.Itoa(start) + rest
}

// return a 11-digit Mainland China hidden middle 4-digit mobile number
func MobileHideMiddle() string {
	start := randomElement(data.MESH_NUM)
	var rest string
	for i := 0; i < 4; i++ {
		random, _ := randomInteger(0, 9)
		rest += strconv.Itoa(random)
	}
	return strconv.Itoa(start) + "****" + rest
}

// return a random Chinese Mainland landline number
func Landline() string {
	var num string
	boolean := randomElement([]bool{true, false})
	areaNum := randomElement(data.TEL_AREA_NUM)
	if boolean {
		random, _ := IntBy(7, true)
		num = areaNum + "-" + strconv.Itoa(random)
	} else {
		random, _ := IntBy(8, true)
		num = areaNum + "-" + strconv.Itoa(random)
	}
	return num
}

// return a random http method
func HttpMethod() string {
	return randomElement(data.HTTP_METHODS)
}

// return a random http status code
func HttpStatusCode() int {
	randomType := randomElement(data.HTTP_STATUS_TYPES)
	return randomElement(data.HTTP_STATUS_CODE[randomType])
}

// return a random uuid that conforms to the RFC 4122
func UUID() string {
	uuid := make([]byte, 16)
	_, err := cryptoRand.Read(uuid)
	if err != nil {
		return ""
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

// generate a random sub-directory string at special sub level. At least one layer
func subDirectory(sub uint) string {
	var (
		result string
		length uint = 1
	)
	if sub > 0 {
		length = sub
	}
	for i := 0; i < int(length); i++ {
		random, _ := randomInteger(1, 6)
		w := "/" + Word(uint(random), "en")
		result += w
	}

	return result
}
