package algo

import (
	"errors"
)

// A priority queue that can hold any set of items that implements
// the PQItem interface
type PriorityQueue struct {
	// Our internal store of data. To make math easier, we start
	// storing data in data[1]. data[0] is unused.
	data []PQItem

	// The current number of actual items
	n int

	// The max number of items we can store
	maxN int

	// Given a PQItem, its index in our data array
	revMap map[PQItem]int
}

// Implement this interface to use this type on a PriorityQueue
type PQItem interface {
	// Compare the receiver of the method call to item passed in args. The
	// implementing function should use type assertion to convert the passed
	// in item. Return true if receiver should be considered a lower
	// priority than the argument, false otherwise. How you implement
	// this for your item determines whether the priority queue is a
	// "min queue" or a "max queue".
	PQLess(PQItem) bool
}

var NotFound = errors.New("item not found in priority queue")
var PQEmpty = errors.New("priority queue is empty")
var PQFull = errors.New("priority queue is full")

// Create a new priority queue with max size of maxN
func NewPriorityQueue(maxN int) *PriorityQueue {
	return &PriorityQueue{
		data:   make([]PQItem, maxN+1),
		maxN:   maxN,
		revMap: make(map[PQItem]int),
	}
}

// What is the index value of this item?
func (pq *PriorityQueue) IndexOf(key PQItem) (int, error) {
	val, ok := pq.revMap[key]

	if !ok {
		return -1, NotFound
	}

	return val, nil
}

// Do we contain this item?
func (pq *PriorityQueue) Contains(key PQItem) bool {
	_, ok := pq.revMap[key]

	return ok
}

// We have changed the value as used by PQLess() for this key. Sink and
// swim above/below to be sure value in correct place.
func (pq *PriorityQueue) IndicateChange(i int) {
	pq.swim(i)
	pq.sink(i)
}

// Insert a new value into our priority queue
func (pq *PriorityQueue) Insert(key PQItem) error {

	if pq.n >= pq.maxN {
		return PQFull
	}

	pq.n++

	// Start by storing the item at the end of our data array
	pq.revMap[key] = pq.n
	pq.data[pq.n] = key

	// Restore heap order by swimming up to place it in the correct location
	pq.swim(pq.n)

	return nil
}

// Delete this item from the priority queue. Returns NotFound if item
// is not currently in the queue.
func (pq *PriorityQueue) Delete(key PQItem) error {

	// Get the index or return an error
	i, err := pq.IndexOf(key)
	if err != nil {
		return err
	}

	// Swap this value with the last value
	pq.swap(i, pq.n)

	// Remove from reverse map
	delete(pq.revMap, pq.data[pq.n])

	// Delete the swapped value
	pq.data[pq.n] = nil

	// Decrement our total count of objects
	pq.n--

	// Generally, we want to restore heap property from i, but it's possible i
	// was the last index. Get the min of i and pq.n.
	if i > pq.n {
		i = pq.n
	}

	// Restore heap property if we still have any values
	if i > 0 {
		pq.swim(i)
		pq.sink(i)
	}

	return nil
}

// GetMax returns the highest priority value without deleting it
func (pq *PriorityQueue) GetMax() (PQItem, error) {
	if pq.n < 1 {
		return nil, PQEmpty
	}

	return pq.data[1]
}

// Delete the highest priority item in our queue
func (pq *PriorityQueue) DelMax() (PQItem, error) {
	if pq.n < 1 {
		return nil, PQEmpty
	}

	// The max item is index 1, get it
	max := pq.data[1]

	// Put the n-indexed item on top
	pq.swap(1, pq.n)

	// Remove from reverse map
	delete(pq.revMap, pq.data[pq.n])

	// Delete the old swapped max
	pq.data[pq.n] = nil

	// Our queue is one item shorter now
	pq.n--

	// Restore heap order
	pq.sink(1)

	return max, nil
}

// Starting from k, "swim" higher priority values up toward index 1. This
// restores heap order from k back to index 1 when a new value is placed at
// k.
func (pq *PriorityQueue) swim(k int) {
	// While we are not at the top index, and there are still values
	// out of heap order
	for k > 1 && pq.data[k/2].PQLess(pq.data[k]) {

		// Swap k with its left child
		pq.swap(k/2, k)

		// Go up priority tree one level
		k = k / 2
	}
}

// Starting with k, "sink" lower priority queue values down the tree. This
// restores heap down from k down to the end of the queue when a new
// value is inserted at k.
func (pq *PriorityQueue) sink(k int) {
	// While we have a left child of k within the size of our data
	for 2*k <= pq.n {

		// Find the indexes of the left and right children of k
		leftChild := 2 * k
		rightChild := 2*k + 1

		// By default, assume left child is biggest, but if
		// rightChild isn't greater than n, it could also
		// be it.
		j := leftChild

		if rightChild <= pq.n && pq.data[leftChild].PQLess(pq.data[rightChild]) {
			j = rightChild
		}

		// j has the largest child, now check against our parent
		if pq.data[k].PQLess(pq.data[j]) {

			// If the parent is less, we should swap with j
			pq.swap(k, j)
		}

		// Set k for the next round down the tree
		k = j
	}
}

// Swap values at index i and j
func (pq *PriorityQueue) swap(i, j int) {

	// Swap the reverse map
	pq.revMap[pq.data[i]], pq.revMap[pq.data[j]] = pq.revMap[pq.data[j]], pq.revMap[pq.data[i]]

	// Swap the actual data
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]

}

// Is our priority queue empty?
func (pq *PriorityQueue) IsEmpty() bool {
	return pq.n < 1
}

// What is the maximum size we can hold in our priority queue?
func (pq *PriorityQueue) MaxSize() int {
	return pq.maxN
}

// What is the current number of items in the priority queue?
func (pq *PriorityQueue) Size() int {
	return pq.n
}
