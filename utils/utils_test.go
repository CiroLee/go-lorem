package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDecimalPartLen(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	len1 := GetDecimalPartLen(10.12)
	len2 := GetDecimalPartLen(10.00)

	is.Equal(len1, 2)
	is.Equal(len2, 0)
}
