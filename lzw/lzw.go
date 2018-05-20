package lzw

import (
	"bufio"
	"encoding/hex"
	"errors"
	"io"

	"github.com/brnstz/algo/bit"
)

const (
	// Global variables that are true regardless of codewordSize
	byteSize       = 8
	initialCodeMax = 1 << byteSize

	// Our default codewordSize
	defaultCodewordSize = 12
)

var (
	// ErrDecoding is returned when unexpected data is found in the
	// stream we are decoding
	ErrDecoding = errors.New(
		"unexpected bits in stream during decoding",
	)
)

type translations struct {
	// Internal maps of encoded and decoded values. We use a hex string
	// to map encoded bytes to their code because a slice of bytes cannot
	// be a key in a map.
	encoded map[string]int
	decoded map[int][]byte

	// nextCode is the next code we are going to use when adding a translation
	nextCode int

	// codewordSize is the number of bits in a code
	codewordSize int

	// codeBytes is the number of bytes required to fully store a code
	// (even if we don't use all of the bits)
	codeBytes int
}

func newTranslations(codewordSize int) *translations {

	// The number of bytes required to store a code are at least codewordSize /
	// byteSize. If it's not an even byte, then add one.
	codeBytes := codewordSize / byteSize
	if codewordSize%byteSize != 0 {
		codeBytes++
	}

	t := &translations{
		encoded:      createEncodedMap(),
		decoded:      createDecodedMap(),
		codewordSize: codewordSize,
		nextCode:     initialCodeMax + 1,
		allCodeMax:   1 << codewordSize,
		codeBytes:    codeBytes,
	}

	return t
}

// Add creates a translation for this set of decoded bytes. We return the new
// code and a true value that indicates the translation was created. If a
// mapping for this decoded set of bytes already exists, we return that value
// and false as the second value. If there is no room to add a new code, we
// return -1 as the code and false as the second value.
func (t *translations) Add(decoded []byte) (int, bool) {
	var (
		code   int
		exists bool
	)

	// If we don't have any room, then don't add it.
	if t.nextCode >= t.allCodeMax {
		return -1, false
	}

	// If we already have a code, then use that
	code, exists = t.encoded[hex.EncodeToString(decoded)]
	if exists {
		return code, false
	}

	// Get next available code
	code = t.nextCode

	// Set encoded map
	t.encoded[hex.EncodeToString(decoded)] = code
	t.decoded[code] = decoded

	// Increment for next code to add
	t.nextCode++

	return code, true
}

// GetDecoded returns the translated bytes for this code if they exist. If they
// do not exist, false is returned as the second value.
func (t *translations) GetDecoded(code int) ([]byte, bool) {
	return t.decoded[code]
}

// GetEncoded returns the encoded int for this translation if it exists. If
// it does not exists, false is returned as the second value.
func (t *translations) GetEncoded(decoded []byte) (int, bool) {
	return t.encoded[hex.EncodeToString(decoded)]

}

func createEncodedMap() map[string]int {
	m := map[string]int{}

	// Every 8-bit character gets mapped from hex to itself
	for i := 0; i < initialCodeMax; i++ {
		m[hex.EncodeToString([]byte{byte(i)})] = i
	}

	return m
}

func createDecodedMap() map[int][]byte {
	m := map[int][]byte{}

	// Every 8-bit character gets mapped from itself to a list of bytes
	for i := 0; i < initialCodeMax; i++ {
		m[i] = []byte{byte(i)}
	}

	return m
}

// Encode reads uncompressed data from r and writes a compressed version to w
func Encode(r io.Reader, w io.Writer) error {
	var (
		err    error
		b      byte
		buff   []byte
		exists bool
		code   int
		i      uint8
	)

	t := newTranslations(defaultCodewordSize)

	output := make([]byte, codeBytes)
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
		_, exists = t.GetEncoded(buff)
		if exists {
			continue
		}

		// If it didn't exist, then give up and write the code for
		// everything except this current character.
		code = t.GetEncoded(buff[0 : len(buff)-1])

		// Create the output list of bytes
		for i = 0; i < codeBytes; i++ {
			output[i] = byte(code >> (byteSize * i))
		}

		// Write exactly codewordSize bits downstream
		// FIXME: don't use constant here
		err = bitw.WriteBits(output, defaultCodewordSize)
		if err != nil {
			return err
		}

		// If we have room for new codes, then add it
		t.Add(buff)

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
	code = t.Get(buff)

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

func itob(i int) []byte {

}

// Decode reads compressed data from r and writes an uncompressed version to
// w
func Decode(r io.Reader, w io.Writer) error {
	var (
		// buffers of untranslated codewords
		code     []byte
		lastCode []byte

		// buffers of translated bytes
		translation []byte
		newEntry    []byte

		firstCharLastTranslation byte

		err          error
		exists       bool
		backupExists bool
	)

	// Initialize reader/writer and initial code mapping
	bitr := bit.NewReader(r)
	bw := bufio.NewWriter(w)
	t := newTranslations(defaultCodewordSize)

	// Read the first codeword from the incoming stream
	code, _, err = bitr.ReadBits(codewordSize)
	if err != nil {
		break
	}

	// Get the decoded list of bytes of this codeword. If the first code isn't
	// in our initial map, something is wrong.
	// FIXME
	translation, exists = t.Get(decoded[btoi(code)])
	if !exists {
		return ErrDecoding
	}

	// Write the decoded value
	_, err = bw.Write(translation)
	if err != nil {
		return err
	}

	// Save for detecting new translations
	lastCode = code
	firstCharLastTranslation = translation[0]

	// Continue to read every other codeword
	for {

		// Read the next codeword from the incoming stream
		code, _, err = bitr.ReadBits(codewordSize)
		if err != nil {
			break
		}

		// Get the decoded list of bytes of this codeword
		translation, exists = decoded[btoi(code)]
		if !exists {

			// If the current translation doesn't exists, first
			// look to our last code
			translation, backupExists = decoded[btoi(lastCode)]

			// If lastCode doesn't exist, we have a corrupted
			// input file.
			if !backupExists {
				return ErrDecoding
			}

			// We can infer the translation of the current code by taking
			// the translation of lastCode and appending the first
			// character of the last translation
			translation = append(translation, firstCharLastTranslation)
		}

		// Write the decoded value
		_, err = bw.Write(translation)
		if err != nil {
			return err
		}

		// Save for next iteration
		firstCharLastTranslation = translation[0]

		// Save the two character entry that will inform our future
		// translations
		newEntry = []byte{oldCode, firstCharLastTranslation}

		// Save for next iteration
		lastCode = code
	}

	return nil
}
