package lzw

import (
	"bufio"
	"errors"
	"io"

	"github.com/brnstz/algo/bit"
)

const defaultCodewordSize = 12

var (
	// ErrDecoding is returned when unexpected data is found in the
	// stream we are decoding
	ErrDecoding = errors.New(
		"unexpected bits in stream during decoding",
	)

	// ErrEncoding is returned when an unexpected state occurs
	// during encoding
	ErrEncoding = errors.New(
		"unexpected error while encoding",
	)
)

// Encode reads uncompressed data from r and writes a compressed version to w
func Encode(r io.Reader, w io.Writer) error {
	var (
		err    error
		b      byte
		buff   []byte
		exists bool
		code   int
	)

	t := newTranslations(defaultCodewordSize)

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
		code, exists = t.GetEncoded(buff[0 : len(buff)-1])
		if !exists {
			return ErrEncoding
		}

		// Write exactly codewordSize bits downstream
		err = bitw.WriteBits(t.Itob(code), t.CodewordSize)
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
	code, exists = t.GetEncoded(buff)
	if !exists {
		return ErrEncoding
	}

	// Write exactly t.CodewordSize bits downstream
	err = bitw.WriteBits(t.Itob(code), t.CodewordSize)
	if err != nil {
		return err
	}

	// Ensure buffer is flushed
	return bitw.Flush()
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
	code, _, err = bitr.ReadBits(t.CodewordSize)
	if err != nil {
		return err
	}

	// Get the decoded list of bytes of this codeword. If the first code isn't
	// in our initial map, something is wrong.
	translation, exists = t.GetDecoded(t.Btoi(code))
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
		code, _, err = bitr.ReadBits(t.CodewordSize)
		if err != nil {
			break
		}

		// Get the decoded list of bytes of this codeword
		translation, exists = t.GetDecoded(t.Btoi(code))
		if !exists {

			// If the current translation doesn't exists, first
			// look to our last code
			translation, backupExists = t.GetDecoded(t.Btoi(lastCode))

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
		newEntry = append([]byte(nil), lastCode...)
		newEntry = append(newEntry, firstCharLastTranslation)
		t.Add(newEntry)

		// Save for next iteration
		lastCode = code
	}

	if err == io.EOF {
		err = nil
	}

	return nil
}
