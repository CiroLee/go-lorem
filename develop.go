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

func sha(length uint) string {
	str := "0123456789abcdefABCDEF"
	return StrBy(length, str)
}
