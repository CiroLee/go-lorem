package core

import (
	"github.com/CiroLee/go-lorem/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomInteger(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1, err1 := RandomInteger(10, 9)
	r2, err2 := RandomInteger(10, 10)
	r3, err3 := RandomInteger(0, 10)

	is.Equal(r1, 0)
	is.Error(err1)
	is.Equal(r2, 0)
	is.Error(err2)
	is.GreaterOrEqual(r3, 0)
	is.LessOrEqual(r3, 10)
	is.Nil(err3)
}

func TestRandomFloat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1, err1 := RandomFloat(10, 9, 4)
	r2, err2 := RandomFloat(10, 10, 4)
	r3, err3 := RandomFloat(0, 10, 0)
	dLen1 := utils.GetDecimalPartLen(r3)
	r4, err4 := RandomFloat(0, 10, 2)
	dLen2 := utils.GetDecimalPartLen(r4)

	is.Equal(r1, 0.0)
	is.Error(err1)
	is.Equal(r2, 0.0)
	is.Error(err2)
	is.GreaterOrEqual(r3, 0.0)
	is.LessOrEqual(r3, 10.0)
	is.LessOrEqual(dLen1, 4)
	is.Nil(err3)
	is.LessOrEqual(dLen2, 2)
	is.Nil(err4)
}
