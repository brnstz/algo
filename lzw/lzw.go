package lzw

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	initialCodeMax = 1 << 8
	allCodeMax     = 1 << 12
)

func createInitialMap() map[string]int {
	m := map[string]int{}

	// Every 8-bit character gets mapped from hex to itself
	for i := 0; i < initialCodeMax; i++ {
		m[hex.EncodeToString([]byte{byte(i)})] = i
	}

	return m
}

// Encode reads uncompressed data from r and writes it to w
func Encode(r io.Reader, w io.Writer) error {
	var (
		err      error
		b        byte
		buff     []byte
		nextCode = initialCodeMax + 1
		exists   bool

		// outputByte byte
	)

	codes := createInitialMap()
	br := bufio.NewReader(r)
	// bw := bufio.NewWriter(w)

	b, err = br.ReadByte()
	if err != nil {
		return err
	}

	buff = append(buff, b)

	for {

		b, err = br.ReadByte()

		if err != nil {
			break
		}

		buff = append(buff, b)

		// fmt.Println(buff)

		// If current buff is in our code, then continue
		// and try to find a bigger code
		_, exists = codes[hex.EncodeToString(buff)]
		if exists {
			continue
		}

		// Otherwise, print out the code without the most recent
		// byte
		// FIXME: figure out how to align bytes
		fmt.Printf("%v => %v\n",
			buff[0:len(buff)-1],
			codes[hex.EncodeToString(buff[0:len(buff)-1])],
		)

		// FIXME: what to do when we run out of codes?
		if nextCode < allCodeMax {
			codes[hex.EncodeToString(buff)] = nextCode
			nextCode++
		}

		buff = []byte{b}

	}

	// Print final code. Also FIXME to figure out how to align bytes
	fmt.Println(codes[hex.EncodeToString(buff)])

	if err == io.EOF {
		err = nil
	}

	return err
}
