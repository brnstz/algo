package algo_test

import (
	"algo"

	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

// A generic interface for queues and stacks so we can use the same
// test code
type AddDel interface {
	Add(interface{})
	Del() (interface{}, error)
}

func loadIt(ad AddDel, file string, t *testing.T) {

	fh, err := os.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()

	// Read curVal as a string. If it's "-", that means delete. Otherwise
	// it should be an int to add which we translate using Atoi()
	var curVal string

	for {
		// Scan a string into curVal
		_, err = fmt.Fscan(fh, &curVal)

		// If EOF, stop reading file. Other errors are fatal.
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}

		if curVal == "-" {
			// Dash means to delete from our stack/queue
			_, err = ad.Del()
			if err != nil {
				t.Fatal(err)
			}
		} else {
			// Otherwise convert to an int and add to our object.
			intVal, err := strconv.Atoi(curVal)
			if err != nil {
				t.Fatal(err)
			}

			ad.Add(intVal)
		}
	}
}

func TestQueue(t *testing.T) {
	q := algo.NewQueue()

	loadIt(q, "data/numdash.txt", t)

	// Check the expected value of size
	size := q.Size()
	if size != 30714 {
		t.Fatalf("Expected size 30714 but got %v", size)
	}

	// Check expected value on top of queue
	peek := q.Peek()
	if peek != 1270713267 {
		t.Fatalf("Expected 1270713267 on top of queue, but got: %v", peek)
	}

	// Delete all values from queue, ensure that empty works and that down
	// sizing works without a panic.
	for !q.IsEmpty() {
		q.Dequeue()
	}

	size = q.Size()
	if size != 0 {
		t.Fatalf("Expected queue with size 0, but got size: %v", size)
	}
}

func TestStack(t *testing.T) {
	s := algo.NewStack()

	loadIt(s, "data/numdash.txt", t)

	// Check the expected value of size
	size := s.Size()
	if size != 30714 {
		t.Fatalf("Expected size 30714 but got %v", size)
	}

	// Check expected value on top of stack
	peek := s.Peek()
	if peek != 318299769 {
		t.Fatalf("Expected 318299769 on top of queue, but got: %v", peek)
	}

	// Exhuastively pop all values off stack
	for !s.IsEmpty() {
		s.Pop()
	}

	// We should get a nil and error value here
	beNil, popErr := s.Pop()
	if beNil != nil || popErr != algo.EmptyStack {
		t.Fatalf("Expected nil and empty stack error, got %v and %v", beNil, popErr)
	}

	// We should get size 0 here
	size = s.Size()
	if size != 0 {
		t.Fatalf("Expected queue with size 0, but got size: %v", size)
	}
}
