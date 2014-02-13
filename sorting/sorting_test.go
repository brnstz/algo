package sorting_test

import (
	"algo/sorting"

	"fmt"
	"io"
	"os"
	"sort"
	"testing"
)

// Load unsorted data from filename. Return a sort.StringSlice. This can be
// converted to sorting.StringSlice (supports CopyAux()) for mergesort.
func createSortSlice(filename string, t *testing.T) (sort.StringSlice, int) {
	// Create an empty slice of strings, using interface from the sorting
	// package
	strSlc := sort.StringSlice{}

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

// Create a slice with the sorting.StringSlice interface by converting a
// sort.StringSlice
func createSortingSlice(filename string, t *testing.T) (sorting.StringSlice, int) {
	sortSlice, count := createSortSlice(filename, t)
	sortingSlice := make(sorting.StringSlice, count)

	for i, val := range sortSlice {
		sortingSlice[i] = val
	}

	return sortingSlice, count
}

// Verify a sort.StringSlice by checking each successive value is >=
func verifySort(strSlc sort.StringSlice, count int, t *testing.T) {
	// Check that list sort has suceeded
	for i := 0; i < count-1; i++ {
		if strSlc.Less(i+1, i) {
			t.Fatal("These values are not sorted:", strSlc[i:i+2])
		}
	}
}

// Verify a sorting.StringSlice by checking each successive value is >=
func verifySorting(strSlc sorting.StringSlice, count int, t *testing.T) {
	// Check that list sort has suceeded
	for i := 0; i < count-1; i++ {
		if strSlc.Less(i+1, i) {
			t.Fatal("These values are not sorted:", strSlc[i:i+2])
		}
	}
}
func TestInsertionSort(t *testing.T) {
	strSlc, count := createSortSlice("../data/words3.txt", t)

	sorting.InsertionSort(strSlc)

	verifySort(strSlc, count, t)
}

func TestMergeSortTopDown(t *testing.T) {
	strSlc, count := createSortingSlice("../data/words3.txt", t)
	// Make an aux copy of the string slice, so we can use it to
	// temporarily store values as we merge
	aux := make(sorting.StringSlice, count)

	sorting.MergeSortTopDown(strSlc, aux, 0, count-1)
	verifySorting(strSlc, count, t)
}

func TestMergeSortBottomUp(t *testing.T) {
	strSlc, count := createSortingSlice("../data/words3.txt", t)
	// Make an aux copy of the string slice, so we can use it to
	// temporarily store values as we merge
	aux := make(sorting.StringSlice, count)

	sorting.MergeSortTopDown(strSlc, aux, 0, count-1)
	verifySorting(strSlc, count, t)
}

func TestQuicksort(t *testing.T) {
	strSlc, count := createSortSlice("../data/words3.txt", t)

	sorting.Quicksort(strSlc)

	verifySort(strSlc, count, t)
}
