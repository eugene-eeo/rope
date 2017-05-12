package rope

// Rope is an interface that represents either an internal
// or leaf node in the rope.
type Rope interface {
	SplitAt(i int) (Rope, Rope)
	Slice(a, b int) Rope
	Concat(n Rope) Rope
	Rebalance() Rope

	Index(byte) int
	ByteAt(int) byte
	Value() string
	Length() int

	each(fn func(*Leaf))
}
