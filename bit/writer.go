package bit

import (
	"bufio"
	"io"
)

const byteSize = 8

// Writer implements a way to write to a file one bit at a time. This allows
// clients to write unaligned data between calls.
type Writer struct {
	bpos  uint8
	buffb byte
	bw    *bufio.Writer
}

// NewWriter initializes a new Writer that writes bits to w
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		bpos:  0,
		buffb: 0,
		bw:    bufio.NewWriter(w),
	}
}

// Flush writes any unaligned data and flushes the underlying writer. If
// the number of bits written is not a multiple of 8, the last bits with
// be 0s.
func (w *Writer) Flush() error {
	var err error

	if w.bpos > 0 {
		err = w.bw.WriteByte(w.buffb)
		if err != nil {
			return err
		}
	}

	return w.bw.Flush()
}

// WriteBit writes exactly one bit and sets up internal state so the next call
// will append to the same byte or start the next byte where appropriate. Bytes
// are written to the underlying writer as aligned data is accumulated.
func (w *Writer) WriteBit(bit bool) error {
	var err error

	if bit {
		w.buffb |= 1 << w.bpos
	}

	w.bpos++

	if w.bpos%byteSize == 0 {
		err = w.bw.WriteByte(w.buffb)
		if err != nil {
			return err
		}

		w.bpos = 0
		w.buffb = 0
	}

	return nil
}

// WriteBits writes only the number of bits specified given the data in p.
// That is, if p contains n bytes, but bits is less then n*8, the next
// call to WriteBits will start writing at the unaligned portion of
func (w *Writer) WriteBits(p []byte, bits int) error {
	var (
		bitLen uint8
		bit    bool
		j      uint8
		err    error
	)

	// Iterate over every byte in p
	for _, b := range p {

		if bits < byteSize {
			bitLen = uint8(bits)
		} else {
			bitLen = byteSize
		}

		// Iterate over every applicable bit in b
		for j = 0; j < bitLen; j++ {

			if b&(1<<j) == 1 {
				bit = true
			} else {
				bit = false
			}

			err = w.WriteBit(bit)
			if err != nil {
				return err
			}
		}

		bits -= byteSize
	}

	return nil
}
