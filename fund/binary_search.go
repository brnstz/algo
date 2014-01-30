package fund

// Use type assertions to assert the passed-in value to be the same
// as your actual type.
type SearchItem interface {

	// Is the current item less than SearchItem?
	Less(SearchItem) bool

	// Is the current item equal to SearchItem?
	Equals(SearchItem) bool
}

// Return the index of of item in the sorted slice values. If item
// doesn't exist, return -1.
func Find(item SearchItem, values []SearchItem) int {
	low := 0
	mid := 0
	top := len(values) - 1

	// Continue running so long as low has not overtaken top value
	for low <= top {

		// Find the midpoint between the current top and low.
		mid = ((top - low) / 2) + low

		// Check out midpoint. Is it the correct value? If so, we're done.
		// Return that index.
		if item.Equals(values[mid]) {
			return mid
		}

		// Otherwise, check if our current item is lesser or greater
		// to determine how we should proceed.
		if item.Less(values[mid]) {

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
