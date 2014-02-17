package algo

import (
	"sort"
)

// Expand sort.Interface to include the ability to swap via aux
type Interface interface {
	// Set self[i] = other[j]
	CopyAux(other Interface, i, j int)

	// Copied from sort.Interface
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Create a local alias of StringSlice from the sort package
type StringSlice sort.StringSlice

// Copy from the other obj into ourself. Use a type assertion to
// ensure that otherInt is really a StringSlice. In an implemention
// for another type, you would assert to that type.
func (self StringSlice) CopyAux(otherInt Interface, i, j int) {
	other := otherInt.(StringSlice)

	self[i] = other[j]
}

// Implement the basic sort.Interface functions with Len(), Less(), and Swap()
func (self StringSlice) Len() int {
	return len(self)
}

func (self StringSlice) Less(i, j int) bool {
	return self[i] < self[j]
}

func (self StringSlice) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

// Given that a[lo1:hi1] and a[lo2:h2] are each sorted individually,
// merge them together into a single sorted list that spans
// a[lo1:hi2]. Use aux to store temporary values.
func merge(a, aux Interface, lo1, hi1, lo2, hi2 int) {
	// Start the low list at lo1
	var i = lo1

	// Start the high list at the midpoint, lo2
	var j = lo2

	// Start the overall list index at lo1
	var k = lo1

	for {

		// If both counters are too high, then we are done
		if i > hi1 && j > hi2 {
			break
		}

		if i > hi1 && j <= hi2 {
			// We have exhauseted values in the low list, because i is too
			// high. So we must take j.
			aux.CopyAux(a, k, j)
			j++
			k++
		} else if i <= hi1 && j > hi2 {
			// We have exhuasted values in the high list, becauese j is too
			// high. So we must take i.
			aux.CopyAux(a, k, i)
			i++
			k++
		} else if a.Less(i, j) {
			// We potentially have values in i or j, so compare them and take
			// the lesser. If this is true, the lower is i.
			aux.CopyAux(a, k, i)
			i++
			k++
		} else {
			// Otherwise, j is the lower value so take that.
			aux.CopyAux(a, k, j)
			j++
			k++
		}

	}

	// Copy things from aux back into the main list
	for x := lo1; x <= hi2; x++ {
		a.CopyAux(aux, x, x)
	}
}

// Merge sort top down, by recursively splitting the list in halves and
// merging when call stack returns with sorted sublists.
func MergeSortTopDown(a, aux Interface, lo, hi int) {
	// Find the midpoint
	mid := lo + ((hi - lo) / 2)

	// Cross over, we are finished
	if lo >= hi {
		return
	}

	// Recursively sort the bottom half
	MergeSortTopDown(a, aux, lo, mid)

	// Recursively sort the top half
	MergeSortTopDown(a, aux, mid+1, hi)

	// Bottom and top halves are sorted individually, merge
	// them together
	merge(a, aux, lo, mid, mid+1, hi)
}

// Merge sort by iterating through the list with different merge sizes,
// and building a fully sorted list as the merge size grows.
func MergeSortBottomUp(a, aux Interface) {
	aLen := a.Len()
	var lo, mid, hi int

	// Outer loop doubles merge size on each iteration, starting at 1
	for size := 1; size < aLen; size = size * 2 {

		// Inner loop goes through full array in increments of
		// the next merge size (i.e., * 2) and runs merge on the
		// sublists.
		for lo = 0; lo < aLen-size; lo += size * 2 {
			// Find midpoint for this iteration
			mid = lo + size - 1

			// Find high point for this iteration, don't let it be
			// beyond the full length of the list
			hi = lo + size*2 - 1
			if hi > aLen-1 {
				hi = aLen - 1
			}

			// Merge the two lists
			merge(a, aux, lo, mid, mid+1, hi)
		}
	}
}
