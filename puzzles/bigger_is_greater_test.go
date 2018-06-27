package puzzles

import "testing"

func TestBiggerIsGreater1(t *testing.T) {
	expected := "ba"
	x := BiggerIsGreater("ab")
	if x != expected {
		t.Fatalf("expected %v but got %v", expected, x)
	}
}
