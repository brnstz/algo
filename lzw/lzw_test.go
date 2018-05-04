package lzw_test

import (
	"os"
	"testing"

	"github.com/brnstz/algo/lzw"
)

func TestLZW(t *testing.T) {
	var err error

	r, err := os.Open("../data/tale.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	lzw.Encode(r, os.Stdout)
}
