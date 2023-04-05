package lorem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha(t *testing.T) {
	is := assert.New(t)
	sha := sha(40)
	is.Len(sha, 40)
}

func TestGitCommitSha(t *testing.T) {
	is := assert.New(t)
	sha := GitCommitSha()
	is.Len(sha, 40)
}

func TestGitCommitShortSha(t *testing.T) {
	is := assert.New(t)
	shortSha := GitCommitShortSha()
	is.Len(shortSha, 7)
}

func TestGitBranch(t *testing.T) {
	is := assert.New(t)
	branch := GitBranch()
	is.Contains(branch, "/")
}

func TestVersion(t *testing.T) {
	is := assert.New(t)
	version := Version()
	is.Contains(version, ".")
}
