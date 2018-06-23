package algo

import (
	"testing"
)

func TestGridPathEmpty(t *testing.T) {
	grid := [][]bool{
		{true, false, false},
		{false, true, false},
		{false, false, true},
	}

	// The only possible path is diagonal so it should be nil
	path := GridPath(grid, 0, 0, 2, 2)
	if path != nil {
		t.Fatalf("expected nil path but got: %v", path)
	}
}

func TestGridPathImpossible(t *testing.T) {
	grid := [][]bool{
		{true, false, false},
		{true, true, false},
		{true, true, true},
		{true, false, true},
		{true, false, true},
		{true, false, true},
		{true, false, true},
		{true, false, true},
		{true, false, true},
		{true, true, true},
	}

	// 10,2 isn't on the grid so it should be a null path
	path := GridPath(grid, 0, 0, 10, 2)
	if path != nil {
		t.Fatalf("expected nil path but got: %v", path)
	}
}

func TestGridPathLine(t *testing.T) {
	grid := [][]bool{
		{true, true, true, true, true, true},
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
		{true, true, true, true, true, true},
	}

	path := GridPath(grid, 0, 0, 0, 5)
	if path.Weight != 5 {
		t.Fatalf("expected path of weight 5 but got: %v", path.Weight)
	}

	path = GridPath(grid, 7, 0, 7, 5)
	if path.Weight != 5 {
		t.Fatalf("expected path of weight 5 but got: %v", path.Weight)
	}
}

func TestGridPathMaze(t *testing.T) {
	grid := [][]bool{
		{true, true, true, true, true, true},
		{true, false, false, false, false, false},
		{true, false, false, false, false, false},
		{true, false, false, false, false, false},
		{true, true, true, true, false, false},
		{true, false, false, true, false, false},
		{true, false, false, true, false, false},
		{true, true, true, true, true, true},
	}

	path := GridPath(grid, 0, 0, 7, 3)

	if path.Weight != 10 {
		t.Fatalf("expected path of weight 10 but got: %v", path.Weight)
	}
}
