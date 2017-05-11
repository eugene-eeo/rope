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
	l := rope.L(s.Value())
	assert.Equal(
		t,
		"bcd",
		s.Slice(1, 4).Value(),
	)
	assert.Equal(
		t,
		"bcd",
		l.Slice(1, 4).Value(),
	)
}

func TestConcatNode(t *testing.T) {
	s := rope.Concat(
		rope.L("abc"),
		rope.L("def"),
	)
	n := s.Concat(rope.L("ghi"))
	assert.Equal(t, "abcdefghi", n.Value())
}

func TestByteAt(t *testing.T) {
	s := rope.Concat(
		rope.L("abc"),
		rope.L("def"),
	)
	for i, c := range s.Value() {
		assert.Equal(t, byte(c), s.ByteAt(i))
	}
}

func TestRebalance(t *testing.T) {
	s := rope.Concat(
		rope.L("a"),
		rope.Concat(
			rope.L("b"),
			rope.Concat(
				rope.L("c"),
				rope.Concat(
					rope.L("d"),
					rope.L("e"),
				),
			),
		),
	)
	assert.Equal(t, s.Value(), "abcde")
	assert.Equal(t, s.Rebalance().Value(), "abcde")
}
