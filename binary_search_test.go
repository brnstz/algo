package algo_test

import (
	"github.com/brnstz/algo"

	"fmt"
	"io"
	"os"
	"testing"
)

type IntSlice []int

// Implement the SearchItem interface for ints.
func (values IntSlice) Less(valIn interface{}, i int) bool {
	x := valIn.(int)
	return x < values[i]
}

func (values IntSlice) Equals(valIn interface{}, i int) bool {
	x := valIn.(int)
	return x == values[i]
}

func (values IntSlice) Len() int {
	return len(values)
}

func TestBinarySearch(t *testing.T) {
	values := IntSlice{5, 100, 3422, 9000, 53535}

	if algo.BinarySearch(100, values) != 1 {
		t.Fatal("Expected to find 100 in index 1 of binary search")
	}

	if algo.BinarySearch(50, values) != -1 {
		t.Fatal("Expected to not find 50.")
	}
}

type IntExpectIndex struct {
	Val, Index int
}

func TestBinarySearchViaFile(t *testing.T) {
	fh, err := os.Open("data/sorted-ints.txt")

	if err != nil {
		t.Fatal(err)
	}

	// We know our file has exaactly 1 million records, so we can
	// declare exact size
	values := make(IntSlice, 1000000)
	var i, value int

	// Read from stdin until err is non-nil. This is either EOF
	// or some other error.
	for {

		// Read one int from stdin. Ignore first return value,
		// which will be 1 for number of items read on success.
		_, err = fmt.Fscan(fh, &value)

		if err == nil {
			values[i] = value
			i++

		} else if err == io.EOF {
			break

		} else {
			t.Fatal(err)
		}

	}

	// If the error was not EOF, then we have an unexpected error
	if err != io.EOF {
		t.Fatal(err)
	}

	// Declare a set of tests to run. First part is the value to
	// search for. The second part is the expected index that the
	// binary search will return.
	tests := []IntExpectIndex{
		{500733, 500000},
		{656695, 656546},
		{0, 0},
		{999999, 999999},
		{13, -1},
		{999954, 999965},
		{253528, 252619},
	}

	for _, test := range tests {
		actualIndex := algo.BinarySearch(test.Val, values)
		if test.Index != actualIndex {
			t.Fatalf("Could not find %v at %v, actual index: %v",
				test.Val, test.Index, actualIndex,
			)
		}
	}
}
