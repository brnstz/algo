package algo_test

import (
	"github.com/brnstz/algo"

	"math"
	"math/rand"
	"testing"
	"time"
)

type PQInt int

func (self PQInt) PQLess(otherItem algo.PQItem) bool {
	other := otherItem.(PQInt)
	return self < other
}

// Test priority queue by adding/deleting 100k random entries
func TestPriorityQueueInsertDelete(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	numVals := 100000

	pq := algo.NewPriorityQueue(numVals)

	var val, lastVal PQInt
	for i := 0; i < numVals; i++ {
		val = PQInt(rand.Int31())
		pq.Insert(val)
	}

	size := pq.Size()
	if size != numVals {
		t.Fatalf("Expected exactly %v vals, but found %v", numVals, size)
	}

	lastVal = math.MaxInt32
	for !pq.IsEmpty() {
		valItem, err := pq.DelMax()
		if err != nil {
			t.Fatal(err)
		}
		val = valItem.(PQInt)

		if val > lastVal {
			t.Fatalf("Expected ordering from highest int to lowest int, but found sequenece %v, %v", lastVal, val)
		}
	}
}

type JukeboxSong struct {
	name     string
	priority int
}

// A lower integer value of priority is of higher importance, so use
// > here.
func (self *JukeboxSong) PQLess(otherPQ algo.PQItem) bool {
	other := otherPQ.(*JukeboxSong)

	return self.priority > other.priority
}

// Test other priority queue functionality with a smaller queue
func TestPriorityQueueOther(t *testing.T) {
	item1 := &JukeboxSong{"Thriller", 9000}
	item2 := &JukeboxSong{"Bad Romance", 600}
	item3 := &JukeboxSong{"1999", 200}
	item4 := &JukeboxSong{"Like a Rolling Stone", 10}
	item5 := &JukeboxSong{"Welcome to the Jungle", 30}
	item6 := &JukeboxSong{"Highway to Hell", 800}

	pq := algo.NewPriorityQueue(5)
	pq.Insert(item1)
	pq.Insert(item2)
	pq.Insert(item3)
	pq.Insert(item4)
	pq.Insert(item5)

	err := pq.Insert(item6)
	if err != algo.PQFull {
		t.Fatal("Expected full queue")
	}

	// Top song on initial insert should be item4
	topPQ, err := pq.DelMax()
	if err != nil {
		t.Fatal(err)
	}
	top := topPQ.(*JukeboxSong)
	if top != item4 {
		t.Fatalf("Expected %+v as top item, but got %+v", item4, top)
	}

	// Change the top song to item3
	index, err := pq.IndexOf(item3)
	if err != nil {
		t.Fatal(err)
	}
	item3.priority = 5
	pq.IndicateChange(index)
	topPQ, err = pq.DelMax()
	if err != nil {
		t.Fatal(err)
	}
	top = topPQ.(*JukeboxSong)
	if top != item3 {
		t.Fatalf("Expected %+v as top item now, but got %+v", item3, top)
	}

	// Now delete remaining items one by one
	pq.Delete(item3)
	pq.Delete(item2)
	pq.Delete(item1)
	pq.Delete(item5)

	if pq.Size() != 0 && pq.IsEmpty() {
		t.Fatal("Expected empty priority queue")
	}
}
