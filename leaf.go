package rope

// Leaf is a leaf node.
type Leaf struct {
	value string
}

// L creates a new Leaf node with the given `value`.
func L(value string) *Leaf {
	return &Leaf{value}
}

// SplitAt is similar to Node.SplitAt.
func (l *Leaf) SplitAt(i int) (Rope, Rope) {
	return &Leaf{l.value[:i]}, &Leaf{l.value[i:]}
}

// ByteAt is similar to Node.ByteAt.
func (l *Leaf) ByteAt(i int) byte {
	return l.value[i]
}

// Slice is similar to Node.Slice.
func (l *Leaf) Slice(a, b int) Rope {
	return &Leaf{l.value[a:b]}
}

// Concat is similar to Node.Concat.
func (l *Leaf) Concat(n Rope) Rope {
	return Concat(l, n)
}

// Value is similar to Node.Value.
func (l *Leaf) Value() string {
	return l.value
}

// Length is similar to Node.Length.
func (l *Leaf) Length() int {
	return len(l.value)
}

// Rebalance on a Leaf node does nothing.
func (l *Leaf) Rebalance() Rope {
	return l
}

func (l *Leaf) visit(fn func(*Leaf)) {
	fn(l)
}
