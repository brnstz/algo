package lzw

import (
	"bufio"
	"encoding/hex"
	"io"

	"github.com/brnstz/algo/bit"
)

const (
	codewordSize = 12
	codeBytes    = 2
	byteSize     = 8

	// Initial code (single byte) is just 8 bits
	initialCodeMax = 1 << byteSize

	// Max code size is 12 bits
	allCodeMax = 1 << codewordSize
)

func createInitialMap() map[string]int {
	m := map[string]int{}

	// Every 8-bit character gets mapped from hex to itself
	for i := 0; i < initialCodeMax; i++ {
		m[hex.EncodeToString([]byte{byte(i)})] = i
	}

	return m
}

// Encode reads uncompressed data from r and writes a compressed vesrion to w
func Encode(r io.Reader, w io.Writer) error {
	var (
		err    error
		b      byte
		buff   []byte
		exists bool
		code   int
		i      uint8

		nextCode = initialCodeMax + 1
	)

	output := make([]byte, codeBytes)
	codes := createInitialMap()
	br := bufio.NewReader(r)
	bitw := bit.NewWriter(w)

	// Read the first byte and append to our buffer
	b, err = br.ReadByte()
	if err != nil {
		return err
	}
	buff = append(buff, b)

	for {

		// Peek at the next byte and append to our buffer
		b, err = br.ReadByte()
		if err != nil {
			break
		}
		buff = append(buff, b)

		// If current buff is in our code, then continue and try to find a
		// bigger code
		_, exists = codes[hex.EncodeToString(buff)]
		if exists {
			continue
		}

		// If it didn't exist, then give up and write the code for
		// everything except this current character.
		code = codes[hex.EncodeToString(buff[0:len(buff)-1])]

		for i = 0; i < codeBytes; i++ {
			output[i] = byte(code >> (byteSize * i))
		}

		err = bitw.WriteBits(output, codewordSize)
		if err != nil {
			return err
		}

		// If we have room for new codes, then add it
		if nextCode < allCodeMax {
			codes[hex.EncodeToString(buff)] = nextCode
			nextCode++
		}

		buff = []byte{b}
	}

	// Check any reader error
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		return err
	}

	code = codes[hex.EncodeToString(buff)]

	for i = 0; i < codeBytes; i++ {
		output[i] = byte(code >> (byteSize * i))
	}

	err = bitw.WriteBits(output, codewordSize)
	if err != nil {
		return err
	}

	return bitw.Flush()
}

// Decode reads compressed data from r and writes an uncompressed version to
// w
func Decode(r io.Reader, w io.Writer) error {

}
