package puzzles

import "testing"

func TestQueenAttack1(t *testing.T) {
	x := QueenAttack(4, 0, 4, 4, nil)
	if x != 9 {
		t.Fatalf("expected 9 moves but got %v", x)
	}
}

func TestQueenAttack2(t *testing.T) {
	x := QueenAttack(5, 3, 4, 3, [][]int{{5, 5}, {4, 2}, {2, 3}})
	if x != 10 {
		t.Fatalf("expected 10 moves but got %v", x)
	}
}

func TestQueenAttack3(t *testing.T) {
	x := QueenAttack(1, 0, 1, 1, nil)
	if x != 0 {
		t.Fatalf("expected 0 moves but got %v", x)
	}
}

func TestQueenAttack4(t *testing.T) {
	x := QueenAttack(100000, 0, 4187, 5068, nil)
	if x != 308369 {
		t.Fatalf("expected 308369 moves but got %v", x)
	}
}

func TestQueenAttackNoGrid1(t *testing.T) {
	x := QueenAttack(4, 0, 4, 4, nil)
	if x != 9 {
		t.Fatalf("expected 9 moves but got %v", x)
	}
}

func TestQueenAttackNoGrid2(t *testing.T) {
	x := QueenAttackNoGrid(5, 3, 4, 3, [][]int{{5, 5}, {4, 2}, {2, 3}})
	if x != 10 {
		t.Fatalf("expected 10 moves but got %v", x)
	}
}

func TestQueenAttackNoGrid3(t *testing.T) {
	x := QueenAttackNoGrid(1, 0, 1, 1, nil)
	if x != 0 {
		t.Fatalf("expected 0 moves but got %v", x)
	}
}

func TestQueenAttackNoGrid4(t *testing.T) {
	x := QueenAttackNoGrid(100000, 0, 4187, 5068, nil)
	if x != 308369 {
		t.Fatalf("expected 308369 moves but got %v", x)
	}
}
