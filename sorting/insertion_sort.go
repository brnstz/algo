package sorting

import (
	// Import the core sort package to use its interface, which declares
	// Len(), Less(), and Swap()
	"sort"
)

func InsertionSort(a sort.Interface) {
	// Store the length of the incoming list to be sorteda
	aLen := a.Len()

	// Declare two index variables
	var i, j int

	// Starting with second value, iterate outer loop over the
	// remainder of the list
	for i = 1; i < aLen; i++ {

		// For every value of a[i] from the outer loop, start with index j=i,
		// and go backwards in the list, swapping j with j-1 until you find
		// a lower value in j-1.
		for j = i; j > 0 && a.Less(j, j-1); j-- {

			// Swap lower value back in the list
			a.Swap(j, j-1)
		}
	}
}
