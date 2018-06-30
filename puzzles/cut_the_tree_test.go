package puzzles

import "testing"

func TestCutTree1(t *testing.T) {
	x := CutTree(
		[]int{100, 200, 100, 500, 100, 600},
		[][]int{{1, 2}, {2, 3}, {2, 5}, {4, 5}, {5, 6}})

	t.Fatal(x)
}

func TestCutTree2(t *testing.T) {
	x := CutTree(
		[]int{1, 2, 3, 4, 5, 6},
		[][]int{{1, 2}, {1, 3}, {2, 6}, {3, 4}, {3, 5}})

	t.Fatal(x)
}
