package puzzles

import "testing"

func TestMaxProfit(t *testing.T) {
	var profit int

	profit = MaxProfit([]int{10, 100, 7, 200, 5, 10})
	if profit != 193 {
		t.Fatalf("expected profit of 193 but got %v", profit)
	}

	profit = MaxProfit([]int{1, 1000, 100, 7, 200, 5, 10})
	if profit != 999 {
		t.Fatalf("expected profit of 999 but got %v", profit)
	}

	profit = MaxProfit([]int{1, 1000, 100, 7, 200, 5, 10, 50000, 1, 34, 234, 99, 40000})
	if profit != 49999 {
		t.Fatalf("expected profit of 49999 but got %v", profit)
	}

	profit = MaxProfit([]int{50000, 1000, 100, 7, 200, 5, 10, 50000, 34, 34, 234, 99, 40000, 2000005})
	if profit != 2000000 {
		t.Fatalf("expected profit of 2000000 but got %v", profit)
	}
}
