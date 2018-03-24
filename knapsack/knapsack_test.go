package knapsack_test

import (
	"testing"

	"github.com/brnstz/algo/knapsack"
)

func TestKnapsack(t *testing.T) {
	items := []knapsack.Item{
		knapsack.Item{Weight: 2, Value: 1},
		knapsack.Item{Weight: 100, Value: 100},
		knapsack.Item{Weight: 8, Value: 32},
		knapsack.Item{Weight: 5, Value: 70},
	}

	solution := knapsack.Bounded(items, 50)
	if solution.Value != 103 {
		t.Fatalf("expected 103 but got %v", solution.Value)
	}
}
