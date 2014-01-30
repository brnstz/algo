package fund_test

import (
	"algo/fund"
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

	if fund.Find(100, values) != 1 {
		t.Fatal("Expected to find 100 in index 1 of binary search")
	}

	if fund.Find(50, values) != -1 {
		t.Fatal("Expected to not find 50.")
	}
}
