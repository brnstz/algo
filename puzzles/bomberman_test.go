package puzzles

import (
	"fmt"
	"testing"
)

func TestBomberman(t *testing.T) {
	x := Bomberman(15, []string{
		".......",
		"...O...",
		"....O..",
		".......",
		"OO.....",
		"OO.....",
	})

	expected := []string{
		"OOO.OOO",
		"OO...OO",
		"OOO...O",
		"..OO.OO",
		"...OOOO",
		"...OOOO",
	}

	failed := false
	for i := range x {
		if x[i] != expected[i] {
			failed = true
			break
		}
	}

	if failed {
		fmt.Println("expected:")
		for i := range expected {
			fmt.Println(expected[i])
		}
		fmt.Println()
		fmt.Println("but got:")
		for i := range x {
			fmt.Println(x[i])
		}
		t.Fatal()
	}
}
