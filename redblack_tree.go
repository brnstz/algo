package algo

import (
	"bytes"
	"fmt"
)

// Color of the node, either red or black. Use bool to define so we
// can easily negate the opposite color. (e.g., newColor = !oldColor)
type color bool

const (
	red   color = true
	black color = false
)

// Implement this type to store into Node.Value
type NodeValue interface {
	Less(NodeValue) bool
	Equals(NodeValue) bool
}

// Pointer to root of the tree
type RedBlackTree struct {
	Root *Node
}

// A node on the tree
type Node struct {
	// Node to the left
	Left *Node

	// Node to the right
	Right *Node

	// Color of this node
	Color color

	// Number of nodes under this one
	NodeCount int

	// If we get a dupe node, increment this count
	ValueCount int

	// Value this node holds
	Value NodeValue
}

// Create a red node with this value
func NewNode(v NodeValue) *Node {
	return &Node{
		Value:      v,
		Color:      red,
		NodeCount:  1,
		ValueCount: 1,
	}
}

// Put a new value into a RedBlackTree
func (t *RedBlackTree) Put(v NodeValue) {
	// Send the root and new value to recursive put method. Get
	// the new root in return (may be same, may be different)
	t.Root = t.put(t.Root, v)

	// Root is always black
	t.Root.Color = black
}

// Recursively put v into the tree, returning potential replacement
// for n.
func (t *RedBlackTree) put(n *Node, v NodeValue) *Node {

	// If node is nil, we've found a place for replacement.
	if n == nil {
		return NewNode(v)
	}

	// If v and node equal, then just increment count and return
	// same parent.
	if v.Equals(n.Value) {
		n.ValueCount++
		return n
	}

	// Recursively put to either left or right
	if v.Less(n.Value) {
		n.Left = t.put(n.Left, v)
	} else {
		n.Right = t.put(n.Right, v)
	}

	// If n is not left-leaning, rotateLeft to make it so.
	if !n.Left.isRed() && n.Right.isRed() {
		n = t.rotateLeft(n)
	}

	// There might be two reds in a row. Red violation. Rotate right to fix.
	if n.Left.isRed() && n.Left.Left.isRed() {
		n = t.rotateRight(n)
	}

	// Both links here might be red. Another red violation. Flip colors to fix.
	if n.Left.isRed() && n.Right.isRed() {
		t.flipColors(n)
	}

	n.NodeCount = n.Left.nodeCountZeroNil() + n.Right.nodeCountZeroNil() + 1

	return n
}

// Flip colors to fix two red links at same node
func (t *RedBlackTree) flipColors(n *Node) {
	n.Color = !n.Color

	n.Left.Color = !n.Left.Color
	n.Right.Color = !n.Right.Color
}

// Rotate left at node op (old parent)
func (t *RedBlackTree) rotateLeft(op *Node) *Node {
	//              op = old parent
	//              np = new parent
	//              b = black
	//              r - red
	//
	//
	//                           D (op)
	//                       b /   \r
	//             (op.Left) A      F (op.Right)
	//                            /b  \b
	//                            E   G
	//
	//             ROTATE LEFT TRANSFORMS TO
	//
	//                          F (np)
	//                        r/  \b
	//         (op / np.Left) D    G (np.Right)
	//                      b/ \b
	//                      A   E
	//

	// The new parent is the right child
	np := op.Right

	// The original parent's right child now points to the
	// original right child's left child.
	op.Right = np.Left

	// The new parent's left child is the original parent
	np.Left = op

	// New parent copies color from old parent
	np.Color = op.Color

	// Left link is now red
	np.Left.Color = red

	// Return new parent
	return np
}

// Rotate right at op (old parent)
func (t *RedBlackTree) rotateRight(op *Node) *Node {
	//              op = old parent
	//              np = new parent
	//              b = black
	//              r - red
	//
	//
	//                         F (op)
	//                     r /   \ b
	//           (op.Left) C      H (op.Right)
	//                  /b  \b
	//                  A      E
	//
	//
	//           ROTATE RIGHT TRANSFORMS TO
	//
	//                         C (np)
	//                      b/   \r
	//           (np.Left)  A     F (op / np.Right)
	//                          /b  \b
	//                         E     H
	//

	// The new parent is the left child
	np := op.Left

	// The original parent's left child now points to the new parent's
	// right child.
	op.Left = np.Right

	// The new parent's right child is the original parent
	np.Right = op

	// New parent copies color from old parent
	np.Color = op.Color

	// Right link is now red
	np.Right.Color = red

	return np
}

// Can we find value v in the tree?
func (t *RedBlackTree) Find(v NodeValue) bool {
	return t.find(t.Root, v)
}

// Helper function for Find
func (t *RedBlackTree) find(n *Node, v NodeValue) bool {

	// We reached a nil node, which means we can't find our value
	if n == nil {
		return false
	}

	// We found it, return true
	if v.Equals(n.Value) {
		return true
	}

	if v.Less(n.Value) {
		// If value we're trying to find is less, search left
		return t.find(n.Left, v)

	} else {
		// Otherwise search right
		return t.find(n.Right, v)

	}
}

// What is the height of the tree?
func (t *RedBlackTree) Height() int {
	return t.height(t.Root)
}

// Private helper function for tree height
func (t *RedBlackTree) height(n *Node) int {
	// If this node is nil, the count shouldn't count
	// Or should this be zero?
	if n == nil {
		return -1
	}

	// Get left and right counts recursively
	left := t.height(n.Left)
	right := t.height(n.Right)

	// Return the max height plus 1
	if left > right {
		return 1 + left
	} else {
		return 1 + right
	}
}

// Do a breadth first search and return a string with each level of
// the tree on a newline
func (t *RedBlackTree) BFSString() string {
	// Output buffer for our return string
	out := &bytes.Buffer{}

	// see queue.go for implementation
	// mainQ stores current level
	mainQ := NewQueue()
	mainQ.Enqueue(t.Root)

	// levelQ stores only nodes for the next level
	levelQ := NewQueue()

	i := 0
	fmt.Fprintf(out, "%02d: ", i)
	for {
		// Pull from the main queue
		n_, err := mainQ.Dequeue()

		if err == EmptyQueue {
			// the queue is empty, we're finished
			break
		} else if err != nil {
			// unexpected error
			panic(err)
		}

		// The queue returns an interface{} value. Assert it into *Node and
		// print its value and the number of entries
		n := n_.(*Node)
		fmt.Fprintf(out, "%v (%v) ", n.Value, n.ValueCount)

		// Enqueue child nodes to the next level
		if n.Left != nil {
			levelQ.Enqueue(n.Left)
		}
		if n.Right != nil {
			levelQ.Enqueue(n.Right)
		}

		// If mainQ is empty, then we're at a new level. Print a new line
		// and swap queues
		if mainQ.IsEmpty() && !levelQ.IsEmpty() {
			i++
			fmt.Fprintf(out, "\n%02d: ", i)
			levelQ, mainQ = mainQ, levelQ
		}
	}

	fmt.Fprintln(out)
	return out.String()
}

// Private helper function to check that node is defined and red
func (n *Node) isRed() bool {
	if n != nil && n.Color == red {
		return true
	} else {
		return false
	}
}

// Private helper function to increment node count
func (n *Node) nodeCountZeroNil() int {
	if n == nil {
		return 0
	}

	return n.NodeCount
}
