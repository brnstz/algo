package huffman

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/brnstz/algo"
)

const (
	byteSize = 8
	eofChar  = 26
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

	return n.freq >= otherN.freq
}

// Coder is a Huffman encoder/decoder
type Coder struct {
	valueType int
	root      *node
	r         *bufio.Reader
	rs        io.ReadSeeker
	codeTable map[interface{}][]bool
}

// NewCoder creates a Coder instance that reads from r interpreting values as
// either Binary or Rune
func NewCoder(valueType int, r io.ReadSeeker) (Coder, error) {
	var err error

	c := Coder{
		valueType: valueType,
		r:         bufio.NewReader(r),
		rs:        r,
		codeTable: map[interface{}][]bool{},
	}

	freqs, err := c.createFreq()
	if err != nil {
		return c, err
	}

	c.root, err = c.createTree(freqs)

	c.createCodeTable(c.root, nil)

	for k, v := range c.codeTable {
		fmt.Printf("%v %c => {%v}\n", k, k, v)
	}

	fmt.Println()

	return c, err
}

// Encode writes Huffman encoded data to w
func (c Coder) Encode(w io.Writer) error {
	var (
		err  error
		v    interface{}
		enc  []bool
		ok   bool
		b    byte
		bpos uint8
	)

	// Seek to start of file
	c.rs.Seek(0, io.SeekStart)
	c.r.Reset(c.rs)

	bw := bufio.NewWriter(w)

	for {
		// Get the next value
		v, err = c.getNext()

		if err != nil {
			break
		}

		// Find the huffman coding for the value
		enc, ok = c.codeTable[v]
		if !ok {
			return fmt.Errorf("invalid encoding, unable to find char")
		}

		// Print every bit individually
		for _, bit := range enc {

			// Set bit if necessary
			if bit {
				b |= 1 << bpos
			}

			// Next bit, next loop
			bpos++

			// If we're at an even byte, then write and reset
			if bpos%byteSize == 0 {
				bw.WriteByte(b)

				bpos = 0
				b = 0
			}
		}
	}

	// Ignore EOF error
	if err == io.EOF {
		err = nil
	}

	// Write final byte if we didn't end evenly
	if bpos > 0 {
		bw.WriteByte(b)
	}

	return bw.Flush()
}

func (c Coder) Decode(r io.Reader) {
	var (
		err  error
		b    byte
		bpos uint8

		n *node
	)

	br := bufio.NewReader(r)
	n = c.root

	for {
		// Get every byte in the stream
		b, err = br.ReadByte()

		if err != nil {
			break
		}

		for bpos = 0; bpos < byteSize; bpos++ {
			if b&(1<<bpos) == 0 {
				//log.Printf("n: %v go left\n", n.value)
				n = n.left
			} else {
				//log.Printf("n: %v go right\n", n.value)
				n = n.right
			}

			if n == nil {
				log.Fatal("encountered fatal error")
			}

			if n.value == eofChar {
				break
			}

			if n.value != 0 {
				fmt.Printf("%c", n.value)
				n = c.root
			}
		}
	}

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

	// freqs maps each value to the number of times it occurs. Include
	// an ASCII eofChar with low frequency to indicate end of the stream. Not
	// to be confused with io.EOF (which is an error value, not an ASCII code)
	freqs := map[interface{}]int{
		eofChar: 1,
	}

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

// createCodeTable recursively populates the mapping between node values and
// their Huffman code (as represented by a slice of bools)
func (c Coder) createCodeTable(n *node, enc []bool) {
	// If there is not a null value node, then record its code
	if n.value != 0 {
		c.codeTable[n.value] = enc
	}

	// Recurse left and right
	if n.left != nil {
		newEnc := append([]bool(nil), enc...)
		newEnc = append(newEnc, false)
		c.createCodeTable(n.left, newEnc)
	}

	if n.right != nil {
		newEnc := append([]bool(nil), enc...)
		newEnc = append(newEnc, true)
		c.createCodeTable(n.right, newEnc)
	}
}
