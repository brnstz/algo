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
	// Create an empty slice of strings, using interface from the sort
	// package
	strSlc := sort.StringSlice{}

	// Open a file of strings to sort
	fh, err := os.Open("../data/words3.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()

	// Declare variables to temporarily store a word and count the
	// total number of words.
	var str string
	var count int

	for {
		// Read one string from the file
		_, err := fmt.Fscan(fh, &str)
		if err == io.EOF {
			break
		}

		// Add a new string to the slice
		strSlc = append(strSlc, str)
		count++
	}

	sorting.InsertionSort(strSlc)

	// Check that list sort has suceeded
	for i := 0; i < count-1; i++ {
		if !strSlc.Less(i, i+1) {
			t.Fatal("These values are not sorted:", strSlc[i:i+2])
		}
	}

}
