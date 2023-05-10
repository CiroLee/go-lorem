package lorem

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
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

func TestMD5(t *testing.T) {
	t.Run("md4 test", func(t *testing.T) {
		md5 := MD5()
		assert.Equal(t, 32, len(md5))
	})
}

func TestPassword(t *testing.T) {

	t.Run("Low strength password", func(t *testing.T) {
		password := Password(8, "low")
		assert.Equal(t, 8, len(password), "Password length should be 8")
		assert.True(t, strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz"), "Password should contain lowercase letters")
		assert.False(t, strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_+{}[]:;<>,.?/"), "Password should not contain uppercase letters, digits, or special characters")
	})

	t.Run("Medium strength password", func(t *testing.T) {
		password := Password(12, "medium")
		assert.Equal(t, 12, len(password), "Password length should be 12")
		assert.True(t, strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), "Password should contain lowercase letters, uppercase letters, and digits")
		assert.False(t, strings.ContainsAny(password, "~!@#$%^&*()_+{}[]:;<>,.?/"), "Password should not contain special characters")
	})

	t.Run("High strength password", func(t *testing.T) {
		password := Password(16, "high")
		assert.Equal(t, 16, len(password), "Password length should be 16")
		assert.True(t, strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_+{}[]:;<>,.?/"), "Password should contain lowercase letters, uppercase letters, digits, and special characters")
	})
}
