package algo

import "testing"

func TestPointCombinations(t *testing.T) {
	for i := 0; i < 50; i++ {
		x := PointCombinations(i, AllPossiblePoints)
		t.Errorf("point combinations for %v: %v", i, x)
	}

}
