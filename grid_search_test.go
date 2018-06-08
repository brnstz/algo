package algo

import "testing"

func TestGridSearch(t *testing.T) {

	grid := [][]byte{
		{'h', 'e', 'z', 'z', 'z'},
		{'h', 'e', 'l', 'l', 'z'},
		{'z', 'e', 'z', 'z', 'z'},
		{'z', 'e', 'z', 'l', 'z'},
		{'z', 'e', 'z', 'z', 'z'},
	}

	if GridSearch("hello", grid) != false {
		t.Fatalf("expected false for 'hello' search but got true")
	}

	if GridSearch("hellz", grid) != true {
		t.Fatalf("expected true for 'hellz' search but got false")
	}

	if GridSearch("hellzzzzzzzzzzzz", grid) != true {
		t.Fatalf("expected true for 'hellzzzzzzzzzzzz' search but got false")
	}

	if GridSearch("hellzzzzzzzl", grid) != true {
		t.Fatalf("expected true for 'hellzzzzzzzzzzzzl' search but got false")
	}

	if GridSearch("hellzzzzzzzeeeeeeeeeee", grid) != true {
		t.Fatalf("expected true for 'hellzzzzzzzeeeeeeeeeee' search but got false")
	}

	if GridSearch("hellzhell", grid) != false {
		t.Fatalf("expected false for 'hellzhell' search but got true")
	}

}
