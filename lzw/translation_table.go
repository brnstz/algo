package lzw

import "encoding/hex"

const (
	// Global variables that are true regardless of codewordSize
	byteSize       = 8
	initialCodeMax = 1 << byteSize
)

type translations struct {

	// CodewordSize is the number of bits in a code
	CodewordSize int

	// Internal maps of encoded and decoded values. We use a hex string
	// to map encoded bytes to their code because a slice of bytes cannot
	// be a key in a map.
	encoded map[string]int
	decoded map[int][]byte

	// nextCode is the next code we are going to use when adding a translation
	nextCode int

	// allCodeMax is the maximum number of codes we can store
	allCodeMax int

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
		CodewordSize: codewordSize,

		encoded:    createEncodedMap(),
		decoded:    createDecodedMap(),
		nextCode:   initialCodeMax + 1,
		allCodeMax: 1 << uint(codewordSize),
		codeBytes:  codeBytes,
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
	decoded, exists := t.decoded[code]
	return decoded, exists
}

// GetEncoded returns the encoded int for this translation if it exists. If
// it does not exists, false is returned as the second value.
func (t *translations) GetEncoded(decoded []byte) (int, bool) {
	code, exists := t.encoded[hex.EncodeToString(decoded)]
	return code, exists
}

// Btoi converts a list of bytes to an int, assuming a codewordSize that is
// less than the size of int
func (t *translations) Btoi(p []byte) int {
	var (
		q int
		i int
	)

	// Add each byte to our int
	for i = 0; i < t.codeBytes; i++ {
		q += (int(p[i]) << uint((byteSize * i)))
	}

	return q
}

// Itob converts a code to a list of bytes
func (t *translations) Itob(code int) []byte {
	var i int

	output := make([]byte, t.codeBytes)

	for i = 0; i < t.codeBytes; i++ {
		output[i] = byte(code >> uint((byteSize * i)))
	}

	return output
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
