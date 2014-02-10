package sorting_test

import (
	"algo/sorting"

	"fmt"
	"io"
	"os"
	"testing"
)

func loadUnsortedFile(filename string) (sorting.StringSlice, error) {
	// Create an empty slice of strings, using interface from the sort
	// package
	strSlc := sorting.StringSlice{}

	// Open a file of strings to sort
	fh, err := os.Open(filename)
	if err != nil {
		return strSlc, err
	}
	defer fh.Close()

	// Declare variables to temporarily store a word and count the
	// total number of words.
	var str string

	for {
		// Read one string from the file
		_, err := fmt.Fscan(fh, &str)
		if err == io.EOF {
			break
		}

		// Add a new string to the slice and increment count
		strSlc = append(strSlc, str)
	}

	return strSlc, nil

}

/*
func TestInsertionSort(t *testing.T) {
	strSlc, err := loadUnsortedFile("../data/words3.txt")
	if err != nil {
		t.Fatal("Unable to load file: ", err)
	}
	sorting.InsertionSort(strSlc)

	// Check that list sort has suceeded
	if strSlc[0] != "all" {
		t.Fatal("Expected 'all' in first position of sorted list.")
	}

	if strSlc[strSlc.Len()-1] != "zoo" {
		t.Fatal("Expected 'zoo' in final position of sorted list.")
	}
}
*/

func TestMergeSort(t *testing.T) {
	strSlc, err := loadUnsortedFile("../data/words3.txt")
	if err != nil {
		t.Fatal("Unable to load file: ", err)
	}
	aux := make(sorting.StringSlice, strSlc.Len())

	fmt.Println(strSlc)
	sorting.MergeSortTopDown(strSlc, aux, 0, strSlc.Len()-1)
	fmt.Println(strSlc)

}
