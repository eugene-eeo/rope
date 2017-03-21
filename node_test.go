package rope_test

import "testing"
import "github.com/eugene-eeo/rope"
import "github.com/stretchr/testify/assert"

func TestConcat(t *testing.T) {
	r := rope.Concat(
		rope.L("abc"),
		rope.L("def"),
		rope.L("ghi"),
	)
	assert.Equal(t, "abcdefghi", r.Value())
}

func TestSplitAt(t *testing.T) {
	s := rope.Concat(
		rope.L("abc"),
		rope.L("def"),
	)
	// when i == midpoint
	l, r := s.SplitAt(3)
	assert.Equal(t, "abc", l.Value())
	assert.Equal(t, "def", r.Value())
	// i > midpoint
	l, r = s.SplitAt(4)
	assert.Equal(t, "abcd", l.Value())
	assert.Equal(t, "ef", r.Value())
	// i < midpoint
	l, r = s.SplitAt(2)
	assert.Equal(t, "ab", l.Value())
	assert.Equal(t, "cdef", r.Value())
}

func TestSlice(t *testing.T) {
	s := rope.Concat(rope.L("abc"), rope.L("def"))
	assert.Equal(
		t,
		"bcd",
		s.Slice(1, 4).Value(),
	)
}
