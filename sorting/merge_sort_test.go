package sorting_test

import (
	"algo/sorting"

	"fmt"
	"io"
	"os"
	"testing"
)

func createSortingSlice(filename string, t *testing.T) (sorting.StringSlice, int) {
	// Create an empty slice of strings, using interface from the sorting
	// package
	strSlc := sorting.StringSlice{}

	// Open a file of strings to sort
	fh, err := os.Open(filename)
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

		// Add a new string to the slice and increment count
		strSlc = append(strSlc, str)

		count++
	}

	return strSlc, count
}

func verifySort(strSlc sorting.StringSlice, count int, t *testing.T) {
	// Check that list sort has suceeded
	for i := 0; i < count-1; i++ {
		if !strSlc.Less(i, i+1) {
			t.Fatal("These values are not sorted:", strSlc[i:i+2])
		}
	}
}

func TestMergeSortTopDown(t *testing.T) {
	strSlc, count := createSortingSlice("../data/words3.txt", t)
	// Make an aux copy of the string slice, so we can use it to
	// temporarily store values as we merge
	aux := make(sorting.StringSlice, count)

	sorting.MergeSortTopDown(strSlc, aux, 0, count-1)
	verifySort(strSlc, count, t)
}

func TestMergeSortBottomUp(t *testing.T) {
	strSlc, count := createSortingSlice("../data/words3.txt", t)
	// Make an aux copy of the string slice, so we can use it to
	// temporarily store values as we merge
	aux := make(sorting.StringSlice, count)

	sorting.MergeSortTopDown(strSlc, aux, 0, count-1)
	verifySort(strSlc, count, t)
}
