package huffman

import (
	"bufio"
	"errors"
	"io"

	"github.com/brnstz/algo"
)

const (
	byteSize = 8
	eofChar  = 26
)

var (

	// ErrUnsupportedValueType is
	ErrUnsupportedValueType = errors.New(
		"use either Binary or Rune as valueType",
	)

	// ErrUnexpectedValue is returned when encoding a stream and its
	// huffman coding is not found
	ErrUnexpectedValue = errors.New(
		"input stream has a value not found in training set",
	)

	// ErrDecoding is returned when unexpected data is found in the
	// stream we are decoding
	ErrDecoding = errors.New(
		"unexpected bits in stream during decoding",
	)
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
	codeTable map[interface{}][]bool
	eofChar   interface{}
}

// NewCoder creates a new Huffman coder that trains itself by reading
// values from trainer as valueType. That is, we use the trainer as the
// source of value frequency.
func NewCoder(valueType int, trainer io.Reader) (Coder, error) {
	var err error

	// Initialize the coder
	c := Coder{
		valueType: valueType,
		codeTable: map[interface{}][]bool{},
	}

	switch valueType {
	case Rune:
		c.eofChar = rune(eofChar)

	case Binary:
		c.eofChar = byte(eofChar)

	default:
		return c, ErrUnsupportedValueType
	}

	// Get frequency counts
	freqs, err := c.createFreq(trainer)
	if err != nil {
		return c, err
	}

	// Create the encoding tree and table
	c.root, err = c.createTree(freqs)
	c.createCodeTable(c.root, nil)

	return c, err
}

// Encode writes Huffman encoded data to w given the unencoded source r
func (c Coder) Encode(r io.Reader, w io.Writer) error {
	var (
		err  error
		v    interface{}
		enc  []bool
		ok   bool
		b    byte
		bpos uint8
		done bool
	)

	br := bufio.NewReader(r)
	bw := bufio.NewWriter(w)

	for !done {
		// Get the next value
		v, err = c.getNext(br)

		if err == io.EOF {
			done = true
			err = nil
			v = c.eofChar
		}

		if err != nil {
			break
		}

		// Find the huffman coding for the value
		enc, ok = c.codeTable[v]
		if !ok {
			return ErrUnexpectedValue
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

	// Write final byte if we didn't end evenly
	if bpos > 0 {
		bw.WriteByte(b)
	}

	return bw.Flush()
}

// Decode the stream of data r and write it to w
func (c Coder) Decode(r io.Reader, w io.Writer) error {
	var (
		err  error
		b    byte
		bpos uint8

		n *node
	)

	br := bufio.NewReader(r)
	bw := bufio.NewWriter(w)
	n = c.root

	for {
		// Get every byte in the stream
		b, err = br.ReadByte()

		// If there's a non-nil error, stop reading
		if err != nil {
			break
		}

		// Read each bit of byte
		for bpos = 0; bpos < byteSize; bpos++ {

			// If it's 0, go left, otherwise go right
			if b&(1<<bpos) == 0 {
				n = n.left
			} else {
				n = n.right
			}

			// If we reached a nil node, the file is corrupt
			if n == nil {
				return ErrDecoding
			}

			// If it's EOF, we are done
			if n.value == c.eofChar {
				break
			}

			// If the value is zero, keep on going through
			if n.value == 0 {
				continue
			}

			// If it's not zero, we have something to write
			switch c.valueType {

			case Binary:
				err = bw.WriteByte(n.value.(byte))

			case Rune:
				_, err = bw.WriteRune(n.value.(rune))

			default:
				err = ErrUnsupportedValueType
			}

			// Was there any error in our switch?
			if err != nil {
				return err
			}

			// Reset back to root for next byte
			n = c.root
		}
	}

	return bw.Flush()
}

// getNext gets the next value from the stream, depending on the value type
func (c Coder) getNext(br *bufio.Reader) (interface{}, error) {
	var (
		err error
		v   interface{}
	)

	switch c.valueType {

	case Binary:
		v, err = br.ReadByte()

	case Rune:
		v, _, err = br.ReadRune()

	default:
		return nil, ErrUnsupportedValueType

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

	br := bufio.NewReader(r)

	// freqs maps each value to the number of times it occurs. Include
	// an ASCII eofChar with low frequency to indicate end of the stream. Not
	// to be confused with io.EOF (which is an error value, not an ASCII code)
	freqs := map[interface{}]int{
		c.eofChar: 1,
	}

	// Get frequencies of all values
	for err == nil {
		v, err = c.getNext(br)
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
