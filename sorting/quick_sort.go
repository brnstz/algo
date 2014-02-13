package sorting

import (
	"math/rand"
	"sort"
	"time"
)

// Run a simple implementation of Quicksort on an incoming
// sort.Interface value
func Quicksort(a sort.Interface) {

	// Get the size of the list
	n := a.Len()

	// Shuffle it
	shuffle(a, n)

	// Sort it using recursive calls to quick sort index 0 to len-1
	qsort(a, 0, n-1)
}

// Quicksort has n² complexity in the worst case, when the partitions selected
// are all unbalanced. Shuffling the list before sorting makes this case
// unlikely.
func shuffle(a sort.Interface, n int) {

	// Seed the randomizer with current time
	rand.Seed(time.Now().UnixNano())

	// Run the Knuth / Fisher-Yates shuffle algorithm. Start with i
	// as the last index. Decrement by 1 until i=1.
	for i := n - 1; i > 0; i-- {
		// Swap i with a random j, chosen from index [0..i]
		j := rand.Intn(i + 1)
		a.Swap(i, j)
	}
}

// Find a pivot value between lo and hi, and then place values less than
// pivot before it, and values higher than after it.
func partition(a sort.Interface, lo, hi int) int {

	// Choose some value to be the pivot. Just use the lowest index.
	pivot := lo

	// Whenever we insert a value lower than the pivot, we use this as the
	// index. Start at the beginning of the list, and increment by 1 each time
	// it is used.
	lessIndex := lo

	// Save the pivot value in the highest index
	a.Swap(pivot, hi)

	// Iterate from lo to hi index
	for i := lo; i < hi; i++ {
		// If the current is less than the pivot (stored in a[hi]), then
		// swap to the lessIndex, and increment lessIndex.
		if a.Less(i, hi) {
			a.Swap(lessIndex, i)
			lessIndex++
		}
	}

	// Place pivot in final place, which is one index after all of the low
	// values. All high values are already above it.
	a.Swap(lessIndex, hi)

	// Return the final position of our pivot value
	return lessIndex
}

func qsort(a sort.Interface, lo, hi int) {
	// If lo has crossed hi, there is nothing to sort in this sublist
	if lo >= hi {
		return
	}

	// Parition the list into low and high values, and find the pivot index.
	pivot := partition(a, lo, hi)

	// Sort each sublist
	qsort(a, lo, pivot-1)
	qsort(a, pivot+1, hi)
}
