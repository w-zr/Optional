package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func optionalGet[T any](t *testing.T, val T) {
	i := Optional[T]{}
	_, err := i.Get()
	assert.NotNil(t, err)

	v := i.GetOr(*new(T))
	assert.Equal(t, *new(T), v)

	i.Assign(val)
	v, err = i.Get()
	assert.Nil(t, err)
	assert.Equal(t, val, v)
}

func TestOptionalGet(t *testing.T) {
	cases := []any{
		1,
		0.1,
		"abc",
		[]int{1, 2, 3},
		struct {
			a int
			b float64
			c string
		}{1, 0.1, "abc"},
		[]byte{123},
	}

	for _, c := range cases {
		optionalGet(t, c)
	}
}

func TestMap(t *testing.T) {
	i := Convert(1)
	m := func(i int) int { return i + 1 }
	assert.Equal(t, m(i.MustGet()), i.Map(m).MustGet())
}

func TestFlatMap(t *testing.T) {
	zero := Convert(0)
	fm := func(i int) Optional[int] {
		if i == 0 {
			return Optional[int]{}
		}
		return Convert(10 / i)
	}
	_, err := zero.FlatMap(fm).Get()
	assert.NotNil(t, err)
}
