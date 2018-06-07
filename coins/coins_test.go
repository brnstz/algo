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
