package bit

import (
	"bufio"
	"io"
)

// Reader implements a way to read from a file one bit a time. This allows
// clients to read unaligned data between calls.
type Reader struct {
	bpos  uint8
	buffb byte
	br    *bufio.Reader
}

// NewReader initializes a new Reader than can reads bits from r
func NewReader(r io.Reader) *Reader {
	return &Reader{
		bpos:  0,
		buffb: 0,
		br:    bufio.NewReader(r),
	}
}

// ReadBit reads exactly one bit from our reader and returns it as a bool
func (r *Reader) ReadBit() (bool, error) {
	var (
		err error
		bit bool
	)

	// We need to read a byte everytime we're aligned with byteSize
	if bpos%byteSize == 0 {
		r.buffb, err = r.br.ReadByte(r.buffb)
		if err != nil {
			return err
		}
	}

	// Figure out what the bit is
	if r.buffb&(1<<r.bpos) > 0 {
		bit = true
	} else {
		bit = false
	}

	// Increment for next time
	if r.bpos == (byteLen - 1) {
		r.bpos = 0
	} else {
		r.bpos++
	}

	return bit
}

// ReadBits reads up to the number of bits specified. An array of bytes
// is returned with the minimum number of bytes needed to return the
// data. For example, a read of 12 bits would return two bytes. The total
// number of bits read is returned. io.EOF is returned as error if we
// exceed the number of bits requested.
func (r *Reader) ReadBits(bits int) ([]byte, int, error) {
	var (
		numBytes int
		numBits uint8
		j uint8
	)

	// Calculate the bytes required
	numBytes := bits / byteSize
	if bit%byteSize != 0 {
		numBytes++
	}

	// Create an array with that number of bytes
	p := make([]b, numBytes)

	for i := 0; i < numBytes; i++ {
		for j := 0; j < 
	}
}






