package fund

import (
	"errors"
)

var EmptyQueue = errors.New("queue is empty")

// A FIFO queue implementation using an auto-resizing array
type Queue struct {
	// Our current internal queue holder.
	queue []interface{}

	// Two indexes, both start at zero. Each increment by one each call
	// to their respective functions. Both wrap around to 0 when surpassing
	// the physical size of the current queue.
	enqueueIndex int
	dequeueIndex int

	// Physical size of queue
	queueSize int

	// Number of non-nil things on queue
	logicalSize int
}

// Create a queue with this initial size
func NewQueue() *Queue {
	q := Queue{}

	// Make a new internal queue initially of size 1
	q.queue = make([]interface{}, 1)
	q.queueSize = 1

	return &q
}

// Enqueue a value
func (q *Queue) Enqueue(val interface{}) {

	// If our logical size is greater than our physical size we must resize
	// the queue to be bigger. Double it every time we need to resize.
	if q.logicalSize >= q.queueSize {
		q.resize(q.queueSize * 2)
	}

	// Store the value
	q.queue[q.enqueueIndex] = val

	// Increment our index and logical size
	q.enqueueIndex++
	q.logicalSize++

	// If we surpass the physical size of the queue, wrap around to the
	// front
	if q.enqueueIndex >= q.queueSize {
		q.enqueueIndex = 0
	}
}

// Dequeue an inten
func (q *Queue) Dequeue() (interface{}, error) {

	// If queue is empty, return error
	if q.IsEmpty() {
		return nil, EmptyQueue
	}

	// Get the current item at the dequeue index
	item := q.queue[q.dequeueIndex]

	// Remove reference to deleted item
	q.queue[q.dequeueIndex] = nil

	// Dequeue index goes up one
	q.dequeueIndex++

	// Logical size goes down one
	q.logicalSize--

	// Wrap around if needed.
	if q.dequeueIndex >= q.queueSize {
		q.dequeueIndex = 0
	}

	// If our physical size is quadruple the logical size, let's resize
	// down to twice logical size
	if q.queueSize > (q.logicalSize * 4) {
		q.resize(q.logicalSize * 2)
	}

	return item, nil
}

// Peek at the value on top of the queue
func (q *Queue) Peek() interface{} {

	if q.IsEmpty() {
		return nil
	}

	// Get the current item at the dequeue index
	return q.queue[q.dequeueIndex]
}

// When our internal physical queue is too big or too small, resize to
// a new value, by make
func (q *Queue) resize(newSize int) {

	// Silently refuse to resize below size=1
	if newSize < 1 {
		return
	}

	// Make a new queue of the requested size
	newQueue := make([]interface{}, newSize)

	// i is our index into the old queue. Start at dequeue index.
	i := q.dequeueIndex

	// j is our index into NewQueue
	j := 0

	for {

		// If we reach end of the old queue, go back to 0. This handles wrap-
		// around cases. Reaching nil is still our breaking point.
		if i >= q.queueSize {
			i = 0
		}

		// If old queue val is nil, we've reached a gap
		// and therefore can stop copying
		//log.Printf("%v %v", q, i)
		if q.queue[i] == nil {
			break
		}

		// Another breaking point is when j is greater than original queue
		// size. This means we've looped with no nils.
		if j >= q.queueSize {
			break
		}

		// Copy from old queue to new queue
		newQueue[j] = q.queue[i]

		// Increment for next round
		i++
		j++
	}

	// Make this our new official queue
	q.queue = newQueue

	// It is repacked so that entries start at index 0 and end at j
	q.dequeueIndex = 0
	q.enqueueIndex = j

	// The physical size of the queue is the requested size
	q.queueSize = newSize

	// The new logical size of the queue is the number of elements, which is j
	q.logicalSize = j
}

// How many items are on the queue?
func (q *Queue) Size() int {
	return q.logicalSize
}

// Is the queue empty?
func (q *Queue) IsEmpty() bool {
	return q.logicalSize < 1
}

// Implement AddDel interface so we can use queues and stacks with same
// test code
func (q *Queue) Add(val interface{}) {
	q.Enqueue(val)
}
func (q *Queue) Del() (interface{}, error) {
	return q.Dequeue()
}
