package algo

import "testing"

func TestLeaderboard(t *testing.T) {
	t.Fatal(Leaderboard([]int{100, 90, 90, 90, 80, 80, 80, 79}, []int{70, 75, 80, 90, 100, 110, 99}))
}
