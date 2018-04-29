package huffman_test

import (
	"os"
	"testing"

	"github.com/brnstz/algo/huffman"
)

func TestHuffman(t *testing.T) {

	r, err := os.Open("../data/tale.txt")
	if err != nil {
		t.Fatal(err)
	}

	huff, err := huffman.NewCoder(huffman.Rune, r)
	if err != nil {
		t.Fatal(err)
	}

	huff.Encode(os.Stderr)
}
