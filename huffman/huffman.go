package huffman

import (
	"bufio"
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
func (n node) PQLess(other algo.PQItem) bool {
	otherN := other.(node)

	return n.freq > otherN.freq
}

// Coder is a Huffman encoder/decoder
type Coder struct {
	valueType int
	root      node
	r         *bufio.Reader
}

// NewCoder creates a Coder instance that reads from r interpreting values as
// either Binary or Rune
func NewCoder(valueType int, r io.Reader) (Coder, error) {
	var err error

	c := Coder{
		valueType: valueType,
		r:         bufio.NewReader(r),
	}

	freqs, err := c.createFreq(r)
	if err != nil {
		return c, err
	}

	c.root, err = c.buildTree(freqs)

	return c, err
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
func (c Coder) createFreq(r io.Reader) (map[interface{}]int, error) {
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

// buildTree accepts the frequency count and builds our Huffman tree
func (c Coder) buildTree(freqs map[interface{}]int) (node, error) {
	var (
		pqItem         algo.PQItem
		parent, n1, n2 node
		err            error
	)

	pq := algo.NewPriorityQueue(len(freqs))

	// Create a node for each value and put it into a priority queue
	for v, freq := range freqs {
		n := node{
			freq:  freq,
			value: v,
		}

		err = pq.Insert(n)
		if err != nil {
			return parent, err
		}
	}

	for pq.Size() > 1 {
		// While we still have at least two items, take them and merge
		pqItem, err = pq.DelMax()
		if err != nil {
			return parent, err
		}
		n1 = pqItem.(node)

		pqItem, err = pq.DelMax()
		if err != nil {
			return parent, err
		}
		n2 = pqItem.(node)

		parent = node{
			value: 0,
			freq:  n1.freq + n2.freq,
			left:  &n1,
			right: &n2,
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

		parent = pqItem.(node)
	}

	return parent, nil
}
