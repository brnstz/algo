package puzzles

import (
	"fmt"
	"testing"
)

func TestBomberman(t *testing.T) {
	x := Bomberman(3, []string{
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
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			if x[i][j] != expected[i][j] {
				failed = true
			}
		}
	}

	for i, v := range x {
		for j, char := range v {
		}
		fmt.Println()
	}

	for _, v := range x {
		for _, char := range v {
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}

	t.Fatal()
}
