package lzw_test

import (
	"bytes"
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

	encB := &bytes.Buffer{}
	// decB := &bytes.Buffer{}

	err = lzw.Encode(r, encB)
	if err != nil {
		t.Fatal(err)
	}

	/*
		FIXME
		rb := bytes.NewBuffer(encB.Bytes())

		err = lzw.Decode(rb, decB)
		if err != nil {
			t.Fatal(err)
		}
	*/

}
