package lorem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDecimalPartLen(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	len1 := getDecimalPartLen(10.12)
	len2 := getDecimalPartLen(10.00)

	is.Equal(len1, 2)
	is.Equal(len2, 0)
}
