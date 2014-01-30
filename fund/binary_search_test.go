package binary_search_test

import (
	"algo/fund"
	"testing"
)

// Implement the SearchItem interface for ints.
func (item int) Less(other_ SearchItem) bool {
	other := other_.(int)
	return item < other
}

func (item int) Equals(other_ SearchItem) bool {
	other := other_.(int)
	return item == other
}

func TestBinarySearch(t *testing.T) {
	values := []int{5, 100, 3422, 900}

	if fund.Find(100, values) != 1 {
		t.Fatal("Expected to find 100 in index 1 of binary search")
	}
}
