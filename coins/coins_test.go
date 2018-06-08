package coins_test

import (
	"testing"

	"github.com/brnstz/algo/coins"
)

func TestChangeRecursive(t *testing.T) {
	var combos int

	combos = coins.ChangeRecursive([]int{1, 5, 10, 25, 50, 100}, 100)
	if combos != 293 {
		t.Fatalf("Expected 293 combos but got %v", combos)
	}
}

func TestChangeIterative(t *testing.T) {
	var combos int

	combos = coins.ChangeIterative([]int{1, 5, 10, 25, 50, 100}, 100)
	if combos != 293 {
		t.Fatalf("Expected 293 combos but got %v", combos)
	}
}

func TestChangeLimited(t *testing.T) {
	combos := coins.ChangeLimited([]int{34, 57, 66, 100}, nil, 100)
	if len(combos) != 3 {
		t.Fatalf("Expected 3 combos but got %v", len(combos))
	}

	combos = coins.ChangeLimited([]int{50, 50, 25, 25, 25, 25, 25, 100}, nil, 100)
	if len(combos) != 243 {
		t.Fatalf("Expected 243 combos but got %v", len(combos))
	}
}
