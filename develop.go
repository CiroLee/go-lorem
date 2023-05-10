package lorem

import (
	"fmt"

	"github.com/CiroLee/go-lorem/data"
)

// return a random git commit sha
func GitCommitSha() string {
	return sha(40)
}

// return a random git commit short sha
func GitCommitShortSha() string {
	return sha(7)
}

// return a random git branch
func GitBranch() string {
	return randomElement(data.BRANCHES)
}

// return a random software version
func Version() string {
	major, _ := randomInteger(0, 20)
	minor, _ := randomInteger(0, 99)
	patch, _ := randomInteger(0, 999)

	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

// return a random md5
func MD5() string {
	var str = "0123456789abcdef"
	return StrBy(32, str)
}

// return a random password. support specified length and strength
func Password(length uint, strength string) string {
	var base = "abcdefghijklmnopqrstuvwxyz"
	var chars string
	switch strength {
	case "low":
		chars = base
	case "medium":
		chars = base + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	case "high":
		chars = base + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_+{}[]:;<>,.?/"
	}
	return StrBy(length, chars)
}

func sha(length uint) string {
	str := "0123456789abcdefABCDEF"
	return StrBy(length, str)
}
