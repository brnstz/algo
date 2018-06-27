package puzzles

import "testing"

func TestSquareEncryption(t *testing.T) {
	expected := "hlhe eoe ltr"
	x := SquareEncryption("hello there")
	if x != expected {
		t.Fatalf("expected %v but got %v", expected, x)
	}
}
