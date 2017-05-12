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
	s := rope.Concat(rope.L("abc"), rope.L("def"))
	v := s.Value()
	for i := 0; i < s.Length(); i++ {
		l, r := s.SplitAt(i)
		assert.Equal(t, v[:i], l.Value())
		assert.Equal(t, v[i:], r.Value())
	}
}

func TestSlice(t *testing.T) {
	s := rope.Concat(rope.L("abc"), rope.L("def"))
	v := s.Value()
	for upper := 0; upper < s.Length(); upper++ {
		for lower := 0; lower < upper; lower++ {
			assert.Equal(
				t,
				v[lower:upper],
				s.Slice(lower, upper).Value(),
			)
		}
	}
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

func TestIndex(t *testing.T) {
	s := rope.Concat(rope.L("abc"), rope.L("def"))
	for i, character := range s.Value() {
		assert.Equal(
			t,
			s.Index(byte(character)),
			i,
		)
	}
}
