package lorem

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	n, err := Uint(10, 100)
	n2, err2 := Uint(10, 5)

	is.Nil(err)
	is.Equal(reflect.TypeOf(n).Kind(), reflect.Uint)
	is.Error(err2)
	is.Equal(n2, uint(0))
}

func TestInt(t *testing.T) {
	is := assert.New(t)

	n, err := Int(-10, 10)
	n2, err2 := Int(10, -10)

	is.Nil(err)
	is.Equal(reflect.TypeOf(n).Kind(), reflect.Int)
	is.Error(err2)
	is.Equal(n2, 0)
}

func TestIntBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	p1, err1 := IntBy(0, true)
	p2, err2 := IntBy(2, true)
	p3, err3 := IntBy(2, false)
	p4, err4 := IntBy(1, true)

	is.Equal(p1, 0)
	is.Error(err1)
	is.LessOrEqual(p2, 99)
	is.Nil(err2)
	is.GreaterOrEqual(p3, -99)
	is.Nil(err3)
	is.LessOrEqual(p4, 9)
	is.GreaterOrEqual(p4, 0)
	is.NoError(err4)
}

func TestFloat32(t *testing.T) {
	is := assert.New(t)

	n, err := Float32(3.14, 10.14, 4)
	n2, err2 := Float32(10.14, 3.14, 4)

	is.Nil(err)
	is.Equal(reflect.TypeOf(n).Kind(), reflect.Float32)
	is.Error(err2)
	is.Equal(n2, float32(0))
}

func TestFloat64(t *testing.T) {
	is := assert.New(t)

	n, err := Float64(3.14, 10.14, 4)
	n2, err2 := Float64(10.14, 3.14, 4)

	is.Nil(err)
	is.Equal(reflect.TypeOf(n).Kind(), reflect.Float64)
	is.Error(err2)
	is.Equal(n2, float64(0))
}
