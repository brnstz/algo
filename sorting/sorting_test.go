package sorting_test

import (
	"algo/sorting"

	"fmt"
	"io"
	"os"
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	// Open a file of strings to sort
	fh, err := os.Open("../data/words3.txt")
	if err != nil {
		t.Fatal("Unable to open input file")
	}
	defer fh.Close()

	// Declare variables to temporarily store a word and count the
	// total number of words.
	var str string
	var i int

	// Create an empty slice of strings, using interface from the sort
	// package
	strSlc := sort.StringSlice{}

	for {
		// Read one string from the file
		_, err := fmt.Fscan(fh, &str)
		if err == io.EOF {
			break
		}

		// Add a new string to the slice and increment count
		strSlc = append(strSlc, str)
		i++
	}

	sorting.InsertionSort(strSlc)

	// Check that list sort has suceeded
	if strSlc[0] != "all" {
		t.Fatal("Expected 'all' in first position of sorted list.")
	}

	if strSlc[i-1] != "zoo" {
		t.Fatal("Expected 'zoo' in final position of sorted list.")
	}
}
