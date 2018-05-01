package huffman_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/brnstz/algo/huffman"
)

func TestHuffman(t *testing.T) {

	r, err := os.Open("../data/tale.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	huff, err := huffman.NewCoder(huffman.Rune, r)
	if err != nil {
		t.Fatal(err)
	}

	wb := &bytes.Buffer{}
	huff.Encode(wb)

	rb := bytes.NewBuffer(wb.Bytes())

	huff.Decode(rb)
}
