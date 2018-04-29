package huffman

import (
	"bufio"
	"fmt"
	"io"

	"github.com/brnstz/algo"
)

const byteSize = 8

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

	c.createCodeTable(c.root, 0, 1)

	for k, v := range c.codeTable {
		fmt.Printf("%v %c => {%b %v}\n", k, k, v.code, v.bitLen)
	}

	fmt.Println()

	return c, err
}

// Encode writes Huffman encoded data to w
func (c Coder) Encode(w io.Writer) error {
	var (
		err  error
		v    interface{}
		enc  encoding
		ok   bool
		b    byte
		bpos uint
		epos uint
		i    int
	)

	// Seek to start of file
	c.rs.Seek(0, io.SeekStart)
	c.r.Reset(c.rs)

	bw := bufio.NewWriter(w)

	for err == nil {
		// Reset the position we starting on the encoded value to
		// zero
		epos = 0

		// Get the next value
		v, err = c.getNext()

		// Find the huffman coding for the value
		enc, ok = c.codeTable[v]
		if !ok {
			return fmt.Errorf("invalid encoding, unable to find char")
		}

		fmt.Printf("%c %b %v\n", v, enc.code, enc.bitLen)
		continue
		i++

		if i > 5 {
			break
		}

		// Write the encoded value one byte at a time
		for epos < enc.bitLen {

			fmt.Printf("@1 b: %b, enc.code: %b, enc.bitLen: %v, epos: %v, bpos: %v\n", b, enc.code, enc.bitLen, epos, bpos)

			// Clear relevant bits
			b = b & (0xFF << (byteSize - bpos))

			fmt.Printf("@2 b: %b, enc.code: %b, enc.bitLen: %v, epos: %v, bpos: %v\n", b, enc.code, enc.bitLen, epos, bpos)

			// Set new bits from the code
			b = b | byte((enc.code>>epos)<<bpos)

			fmt.Printf("@3 b: %b, enc.code: %b, enc.bitLen: %v, epos: %v, bpos: %v\n", b, enc.code, enc.bitLen, epos, bpos)

			// epos will either be another byte len or we've reached
			// the end of the bitlen
			if epos+byteSize > enc.bitLen {
				epos = enc.bitLen
			} else {
				epos += byteSize
			}
			fmt.Printf("@4 b: %b, enc.code: %b, enc.bitLen: %v, epos: %v, bpos: %v\n", b, enc.code, enc.bitLen, epos, bpos)

			// bpos will either be zero or the remainder of bits left to be set
			// in this byte
			bpos = epos % byteSize

			// Write every time we're at an aligned byte
			if bpos == 0 {
				fmt.Printf("%b\n", b)
				//bw.WriteByte(b)
			}

			fmt.Printf("@5 b: %b, enc.code: %b, enc.bitLen: %v, epos: %v, bpos: %v\n", b, enc.code, enc.bitLen, epos, bpos)

			fmt.Println()
		}

	}

	// Ignore EOF error
	if err == io.EOF {
		err = nil
	}

	return bw.Flush()
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
	//fmt.Printf("%v %c %b %v\n", n.value, n.value, code, bitLen)

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
		c.createCodeTable(n.right, code|1<<(bitLen+1), bitLen+1)
	}
}
