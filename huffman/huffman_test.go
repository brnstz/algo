package huffman_test

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/brnstz/algo/huffman"
)

func TestHuffman(t *testing.T) {
	var err error

	testVal := "can we encode this"

	r, err := os.Open("../data/tale.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	huff, err := huffman.NewCoder(huffman.Rune, r)
	if err != nil {
		t.Fatal(err)
	}

	rb := bytes.NewBufferString(testVal)
	encB := &bytes.Buffer{}
	decB := &bytes.Buffer{}

	err = huff.Encode(rb, encB)
	if err != nil {
		t.Fatal(err)
	}

	rb = bytes.NewBuffer(encB.Bytes())

	err = huff.Decode(rb, decB)
	if err != nil {
		t.Fatal(err)
	}

	if decB.String() != testVal {
		log.Fatalf("expected '%v' but got '%v'", testVal, decB.String())
	}
}
