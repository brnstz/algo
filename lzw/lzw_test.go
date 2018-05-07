package lzw_test

import (
	"bytes"
	"log"
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

	w := &bytes.Buffer{}

	err = lzw.Encode(r, w)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(w.Len())
}