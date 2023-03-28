package lorem

import (
	"testing"

	"github.com/CiroLee/gear/gearstring"
	"github.com/CiroLee/go-lorem/data"
	"github.com/stretchr/testify/assert"
)

func TestVrm(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	v := Vrm()

	first := gearstring.CharAt(v, 0)
	is.Contains(data.RegionAbbr, first)
	is.LessOrEqual(len(v), 11)
	is.GreaterOrEqual(len(v), 8)
}

func TestVin(t *testing.T) {
	is := assert.New(t)
	vin := Vin()
	is.Len(vin, 17)
}

func TestId(t *testing.T) {
	is := assert.New(t)

	id := Id()
	is.Len(id, 18)
}
