package fund

// Create an interface, steal ideas from part of core sort package.
type SearchSlice interface {
	// Return true if val is less than value stored at index i
	Less(val interface{}, i int) bool

	// Return true if val is equal to value at index i
	Equals(val interface{}, i int) bool

	// Return the length of the slice
	Len() int
}

// Return the index of of item in the sorted slice values. If item
// doesn't exist, return -1.
func Find(item interface{}, values SearchSlice) int {
	low := 0
	mid := 0
	top := values.Len() - 1

	// Continue running so long as low has not overtaken top value
	for low <= top {

		// Find the midpoint between the current top and low.
		mid = ((top - low) / 2) + low

		// Check out midpoint. Is it the correct value? If so, we're done.
		// Return that index.
		if values.Equals(item, mid) {
			return mid
		}

		// Otherwise, check if our current item is lesser or greater
		// to determine how we should proceed.
		if values.Less(item, mid) {

			// Our item is less than the midpoint, so next time, we'll check
			// in vals[low..mid-1]
			top = mid - 1
		} else {

			// Our item is greater than the midpoint, so next time, we'll check
			// in vals[mid+1..top]
			low = mid + 1

		}
	}

	// We can't find it. Return -1
	return -1
}
