package huffman

import (
	"bufio"
	"fmt"
	"io"

	"github.com/brnstz/algo"
)

const (
	// Binary reads the incoming stream as 8-bit values
	Binary = iota
	// Rune reads the incoming stream as UTF-8 characters
	Rune
)

type node struct {
	freq  int
	value interface{}

	left  *node
	right *node
}

// PQLess for a node considers the item with larger frequency to have less
// priority. We want to pop them off from least frequest to most.
func (n *node) PQLess(other algo.PQItem) bool {
	otherN := other.(*node)

	return n.freq > otherN.freq
}

// Coder is a Huffman encoder/decoder
type Coder struct {
	valueType int
	root      *node
	r         *bufio.Reader
	rs        io.ReadSeeker
	codeTable map[interface{}]encoding
}

type encoding struct {
	code   uint64
	bitLen uint
}

// NewCoder creates a Coder instance that reads from r interpreting values as
// either Binary or Rune
func NewCoder(valueType int, r io.ReadSeeker) (Coder, error) {
	var err error

	c := Coder{
		valueType: valueType,
		r:         bufio.NewReader(r),
		rs:        r,
		codeTable: map[interface{}]encoding{},
	}

	freqs, err := c.createFreq()
	if err != nil {
		return c, err
	}

	c.root, err = c.createTree(freqs)

	c.createCodeTable(c.root, 0, 0)

	for k, v := range c.codeTable {
		fmt.Printf("%v => {%b %v}\n", k, v.code, v.bitLen)
	}

	return c, err
}

// Encode writes Huffman encoded data to w
func (c Coder) Encode(w io.Writer) error {
	var (
		err error
		v   interface{}
		enc encoding
		ok  bool
	)

	// Seek to start of file
	c.rs.Seek(0, io.SeekStart)
	c.r.Reset(c.rs)

	for err == nil {
		v, err = c.getNext()

		enc, ok = c.codeTable[v]
		if !ok {
			return fmt.Errorf("invalid encoding, unable to find char")
		}

		fmt.Println(enc)
	}

	// Ignore EOF error
	if err == io.EOF {
		err = nil
	}

	return err
}

// getNext gets the next value from the stream, depending on the value type
func (c Coder) getNext() (interface{}, error) {
	var (
		err error
		v   interface{}
	)

	switch c.valueType {

	case Binary:
		v, err = c.r.ReadByte()

	case Rune:
		v, _, err = c.r.ReadRune()

	}

	return v, err
}

// createFreq reads the incoming stream and creates a mapping of values
// to frequency counts
func (c Coder) createFreq() (map[interface{}]int, error) {
	var (
		v   interface{}
		err error
	)

	// freqs maps each value to the number of times it occurs
	freqs := map[interface{}]int{}

	// Get frequencies of all values
	for err == nil {
		v, err = c.getNext()
		freqs[v]++
	}

	// Ignore EOF error
	if err == io.EOF {
		err = nil
	}

	return freqs, err
}

// createTree accepts the frequency count and builds our Huffman tree
func (c Coder) createTree(freqs map[interface{}]int) (*node, error) {
	var (
		pqItem         algo.PQItem
		parent, n1, n2 *node
		err            error
	)

	pq := algo.NewPriorityQueue(len(freqs))

	// Create a node for each value and put it into a priority queue
	for v, freq := range freqs {
		n := &node{
			freq:  freq,
			value: v,
		}

		err = pq.Insert(n)
		if err != nil {
			return parent, err
		}
	}

	// While we still have at least two items, take them and merge
	for pq.Size() > 1 {
		pqItem, err = pq.DelMax()
		if err != nil {
			return parent, err
		}
		n1 = pqItem.(*node)

		pqItem, err = pq.DelMax()
		if err != nil {
			return parent, err
		}
		n2 = pqItem.(*node)

		// Create a null node with each of these children
		// as different paths. Null nodes allow us to be
		// prefix free.
		parent = &node{
			value: 0,
			freq:  n1.freq + n2.freq,
			left:  n1,
			right: n2,
		}

		err = pq.Insert(parent)
		if err != nil {
			return parent, err
		}
	}

	// Add the final, highest priority node as root
	if !pq.IsEmpty() {
		pqItem, err := pq.DelMax()
		if err != nil {
			return parent, err
		}

		parent = pqItem.(*node)
	}

	return parent, nil
}

func (c Coder) createCodeTable(n *node, code uint64, bitLen uint) {
	//fmt.Println(n.value, n.freq, n.left, n.right, code, bitLen)

	// If there is not a null value node, then record its code
	if n.value != 0 {
		c.codeTable[n.value] = encoding{
			code:   code,
			bitLen: bitLen,
		}
	}

	// Recurse left and right
	if n.left != nil {
		// Left means "0" so just increase the bit length
		c.createCodeTable(n.left, code, bitLen+1)
	}

	if n.right != nil {
		// Right means "1" so add a 1 bit at the new bitLen
		c.createCodeTable(n.right, code|1<<bitLen+1, bitLen+1)
	}
}
