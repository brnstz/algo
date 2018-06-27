package puzzles

import "testing"

func TestRectIntersection(t *testing.T) {
	r1 := &Rect{X1: 3, Y1: 2, X2: 7, Y2: 10}
	r2 := &Rect{X1: 1, Y1: 4, X2: 9, Y2: 5}

	expectedR3 := Rect{X1: 3, Y1: 4, X2: 7, Y2: 5}

	r3 := r1.Intersection(r2)

	if expectedR3 != *r3 {
		t.Fatalf("Expected %v but got %v", expectedR3, *r3)
	}

}

func TestRectNoIntersection(t *testing.T) {
	r1 := &Rect{X1: 1, Y1: 3, X2: 2, Y2: 10}
	r2 := &Rect{X1: 5, Y1: 3, X2: 8, Y2: 6}

	r3 := r1.Intersection(r2)

	if r3 != nil {
		t.Fatalf("Expected nil but got %v", r3)
	}

}
