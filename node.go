package rope

// Node is an internal node in the rope.
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

// NewNode creates a new internal node that concatenates
// all of its arguments. At least two nodes, `a` and `b` are
// required. For example:
//  node := NewNode(NewLeaf("abc"), NewLeaf("def"))
//  NewNode(node NewLeaf("ghi")).Value()
//  // => "abcdefghi"
func NewNode(a, b Rope, nodes ...Rope) *Node {
	s := joinNode(a, b)
	for _, n := range nodes {
		s = joinNode(s, n)
	}
	return s
}

// SplitAt splits the node at the given index, and returns
// the left and right pieces. The left piece contains the
// string up to the i-th character and the right contains
// the remainder, i.e. the i-th character up to the end of
// the string.
//  node.SplitAt(3) // => Leaf{"abc"}, Leaf{"def"}
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
// i.e. from the a-th to the (b-1)-th characters.
//  node.Slice(1, 4) // => Node{Leaf{"bc"}, Leaf{"d"}}
func (n *Node) Slice(a, b int) Rope {
	_, r := n.SplitAt(a)
	l, _ := r.SplitAt(b)
	return l
}

func (n *Node) Concat(x Rope) Rope {
	return NewNode(n, x)
}

func (n *Node) ByteAt(i int) byte {
	l := n.left.Length()
	if i < l {
		return n.left.ByteAt(i)
	}
	return n.right.ByteAt(i - l)
}

func (n *Node) Value() string {
	return n.left.Value() + n.right.Value()
}

func (n *Node) Length() int {
	return n.length
}

func (n *Node) visit(fn func(*Leaf)) {
	n.left.visit(fn)
	n.right.visit(fn)
}

type nodeInfo struct {
	node Rope
	size int
}

func (n *Node) Rebalance() Rope {
	S := []nodeInfo{}
	n.visit(func(leaf *Leaf) {
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
			node = NewNode(prev.node, node)
		}
		S = append(S, nodeInfo{node, size})
	})
	// perform one final pass through the nodes to merge
	// any leftover nodes.
	root := S[0].node
	for _, info := range S[1:] {
		root = NewNode(root, info.node)
	}
	return root
}
