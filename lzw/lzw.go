package lzw

import (
	"bufio"
	"encoding/hex"
	"errors"
	"io"
	"log"

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

var (
	// ErrDecoding is returned when unexpected data is found in the
	// stream we are decoding
	ErrDecoding = errors.New(
		"unexpected bits in stream during decoding",
	)
)

func createInitialMap() map[string]int {
	m := map[string]int{}

	// Every 8-bit character gets mapped from hex to itself
	for i := 0; i < initialCodeMax; i++ {
		m[hex.EncodeToString([]byte{byte(i)})] = i
	}

	return m
}

func createReverseMap() map[int][]byte {
	m := map[int][]byte{}

	// Every 8-bit character gets mapped from itself to a list of bytes
	for i := 0; i < initialCodeMax; i++ {
		m[i] = []byte{byte(i)}
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
	encoded := createInitialMap()
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
		_, exists = encoded[hex.EncodeToString(buff)]
		if exists {
			continue
		}

		// If it didn't exist, then give up and write the code for
		// everything except this current character.
		code = encoded[hex.EncodeToString(buff[0:len(buff)-1])]

		// Create the output list of bytes
		for i = 0; i < codeBytes; i++ {
			output[i] = byte(code >> (byteSize * i))
		}

		// Write exactly codewordSize bits downstream
		err = bitw.WriteBits(output, codewordSize)
		if err != nil {
			return err
		}

		// If we have room for new codes, then add it
		if nextCode < allCodeMax {
			encoded[hex.EncodeToString(buff)] = nextCode
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

	// Get the final code in the buffer
	code = encoded[hex.EncodeToString(buff)]

	// Create the output list of bytes
	for i = 0; i < codeBytes; i++ {
		output[i] = byte(code >> (byteSize * i))
	}

	// Write exactly codewordSize bits downstream
	err = bitw.WriteBits(output, codewordSize)
	if err != nil {
		return err
	}

	// Ensure buffer is flushed
	return bitw.Flush()
}

// btoi converts a list of bytes to an int, assuming a codewordSize that is
// less than the size of int
func btoi(p []byte) int {
	var (
		q int
		i uint8
	)

	// Add each byte to our int
	for i = 0; i < codeBytes; i++ {
		q += (int(p[i]) << (byteSize * i))
	}

	return q
}

// Decode reads compressed data from r and writes an uncompressed version to
// w
func Decode(r io.Reader, w io.Writer) error {
	var (
		//buff []byte
		err error
		//code int

		// codeword as bytes
		cwb []byte
		//exists bool
	)

	// Initialize reader/writer and initial code mapping
	bitr := bit.NewReader(r)
	//bw := bufio.NewWriter(w)
	//decoded := createReverseMap()

	for {
		// Read the first codeword from the incoming stream
		cwb, _, err = bitr.ReadBits(codewordSize)
		if err != nil {
			return err
		}
		if err != nil {
			log.Printf("err: %v", err)
			break
		}

		log.Printf("codeword: %v\n", cwb)
	}

	/*
		// Get the decoded list of bytes of this codeword
		bytes, exists = decoded[btoi(cwb)]
		// If the first code isn't in our initial map, something is wrong.
		if !exists {
			return ErrDecoding
		}

		// Write the decoded value
		_, err = bw.Write(bytes)
		if err != nil {
			return err
		}

		for {

			// Read the first codeword from the incoming stream
			cwb, _, err = bitr.ReadBits(codewordSize)
			if err != nil {
				return err
			}

			// Get the decoded list of bytes of this codeword
			bytes, exists = decoded[btoi(cwb)]
			// If the first code isn't in our initial map, something is wrong.

		}
	*/

	return nil

}
