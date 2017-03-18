package rope

// Leaf is a leaf node.
type Leaf struct {
	value string
}

// NewLeaf creates a new Leaf node with the given `value`.
func NewLeaf(value string) *Leaf {
	return &Leaf{value}
}

func (l *Leaf) SplitAt(i int) (Rope, Rope) {
	return &Leaf{l.value[:i]}, &Leaf{l.value[i:]}
}

func (l *Leaf) ByteAt(i int) byte {
	return l.value[i]
}

func (l *Leaf) Slice(a, b int) Rope {
	return &Leaf{l.value[a:b]}
}

func (l *Leaf) Concat(n Rope) Rope {
	return NewNode(l, n)
}

func (l *Leaf) Value() string {
	return l.value
}

func (l *Leaf) Length() int {
	return len(l.value)
}

func (l *Leaf) Rebalance() Rope {
	return l
}

func (l *Leaf) visit(fn func(*Leaf)) {
	fn(l)
}
