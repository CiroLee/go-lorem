package lorem

import (
	"testing"

	"github.com/CiroLee/gear/gearmath"
	"github.com/stretchr/testify/assert"
)

func TestElements(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r1 := Elements(arr, 2)
	r2 := Elements(arr, 1)
	r3 := Elements(arr, 0)

	is.True(gearmath.IsSubset(arr, r1))
	is.True(gearmath.IsSubset(arr, r2))
	is.Empty(r3)
}
