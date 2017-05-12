package rope

// Node is an internal node in the tree. The zero value of this
// type is meaningless, and should NOT be used.
type Node struct {
	left   Rope
	right  Rope
	length int
}

func joinNode(l, r Rope) *Node {
	return &Node{
		left:   l,
		right:  r,
		length: l.Length() + r.Length(),
	}
}

// Concat creates a new rope node that concatenates all
// of its arguments. For example:
//  node := Concat(NewLeaf("abc"), NewLeaf("def"))
//  Concat(node NewLeaf("ghi")).Value() // => "abcdefghi"
func Concat(a, b Rope, nodes ...Rope) *Node {
	s := joinNode(a, b)
	for _, n := range nodes {
		s = joinNode(s, n)
	}
	return s
}

// SplitAt splits the node at the given index, and returns
// the left and right pieces. The left piece contains the
// string up to but not including the i-th character and the
// right contains the remainder, i.e. the i-th character up
// to the end of the string.
//  l, r := node.SplitAt(3)
//  l.Value() // => "abc"
//  r.Value() // => "def"
func (n *Node) SplitAt(i int) (Rope, Rope) {
	m := n.left.Length()
	if i == m {
		return n.left, n.right
	}
	if i < m {
		l, r := n.left.SplitAt(i)
		return l, r.Concat(n.right)
	}
	l, r := n.right.SplitAt(i - m)
	return n.left.Concat(l), r
}

// Slice returns a "substring" of the node from [a, b);
// i.e. starting from a-th to the (b-1)-th characters.
//  sub := node.Slice(1, 4)
//  sub.Value() == "bcd"
func (n *Node) Slice(a, b int) Rope {
	_, r := n.SplitAt(a)
	l, _ := r.SplitAt(b - a)
	return l
}

// Concat concatenates the current node with another Rope.
//  node.Concat("ghi").Value() // => "abcdefghi"
func (n *Node) Concat(x Rope) Rope {
	return Concat(n, x)
}

// Index returns the 0-based index of the given byte.
//  node.Index(byte("a")) == 0
func (n *Node) Index(b byte) int {
	i := n.left.Index(b)
	if i != -1 {
		return i
	}
	j := n.right.Index(b)
	if j != -1 {
		return n.left.Length() + j
	}
	return -1
}

// ByteAt returns the i-th character in the string.
//  node.ByteAt(0) == 'a'
//  node.ByteAt(2) == 'c'
func (n *Node) ByteAt(i int) byte {
	l := n.left.Length()
	if i < l {
		return n.left.ByteAt(i)
	}
	return n.right.ByteAt(i - l)
}

// Value returns the value of the node by concatenating
// all its children.
//  node.Value() == "abcdef"
func (n *Node) Value() string {
	return n.left.Value() + n.right.Value()
}

// Length returns the length of .Value().
//  node.Length() == 6
func (n *Node) Length() int {
	return n.length
}

func (n *Node) each(fn func(*Leaf)) {
	n.left.each(fn)
	n.right.each(fn)
}

type nodeInfo struct {
	node Rope
	size int
}

// Rebalance explicitly rebalances the node and its children.
// After rebalancing, splitting and insertion operations will
// be faster.
func (n *Node) Rebalance() Rope {
	S := []nodeInfo{}
	n.each(func(leaf *Leaf) {
		// For each leaf node:
		// 1. Try to 'reduce' leftwards as much as possible.
		// 2. Only reduce leftwards if both nodes have the
		//    same number of children.
		var node Rope = leaf
		size := 1
		for len(S) > 0 {
			prev := S[len(S)-1]
			if prev.node == node || prev.size != size {
				break
			}
			// pop from the array and merge the two nodes
			S = S[:len(S)-1]
			size += prev.size
			node = Concat(prev.node, node)
		}
		S = append(S, nodeInfo{node, size})
	})
	// perform one final pass through the nodes to merge
	// any leftover nodes.
	root := S[0].node
	for _, info := range S[1:] {
		root = Concat(root, info.node)
	}
	return root
}
