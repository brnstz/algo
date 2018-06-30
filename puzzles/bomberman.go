package puzzles

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	bempty   = '.'
	bbomb    = 'O'
	bdetsecs = 3
)

// https://www.hackerrank.com/challenges/bomber-man/problem

// bdetKey creates a hashable key from an x, y value
func bdetKey(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

// bdetRevKey returns an x, y coordinate based on a key previously
// created in bdetKey
func bdetRevKey(key string) (int, int) {
	var (
		x, y int
	)
	strs := strings.Split(key, ",")
	x, _ = strconv.Atoi(strs[0])
	y, _ = strconv.Atoi(strs[1])

	return x, y
}

// detOneBomb detonates the bomb at the key bombKey and removes itself
// and its neighbors from the detonation list and the grid
func detOneBomb(bombKey string, bombDetSecs map[string]int, grid [][]rune) {
	var (
		bx, by, x, y int
		neighborKey  string
	)

	// The coords of the bomb
	bx, by = bdetRevKey(bombKey)

	// Deteonate our four neighbors and ourselves
	diffs := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {0, 0},
	}

	for _, diff := range diffs {

		// The coords we are detonating
		x = bx + diff[0]
		y = by + diff[1]

		// Make sure it's not off the grid
		if x < 0 || x >= len(grid) {
			continue
		}

		if y < 0 || y >= len(grid[x]) {
			continue
		}

		// Create they key for the neighbor
		neighborKey = bdetKey(x, y)

		// Delete the neighbor.
		delete(bombDetSecs, neighborKey)
		grid[x][y] = bempty
	}
}

// detBombs checks the detonation time of any bombs and detonates or
// decreases their detonation time accordingly.
func detBombs(bombDetSecs map[string]int, grid [][]rune) {
	var (
		bombsToDetonate []string
	)

	for k, v := range bombDetSecs {

		if v > 1 {
			// Decrement detonation time
			bombDetSecs[k]--

		} else if v == 1 {
			// Detonate the bomb and remove it from the map
			delete(bombDetSecs, k)
			bombsToDetonate = append(bombsToDetonate, k)
		}
	}

	for _, v := range bombsToDetonate {
		detOneBomb(v, bombDetSecs, grid)
	}
}

// placeBombs places a bomb in all grid spaces that don't currently
// have one.
func placeBombs(bombDetSecs map[string]int, grid [][]rune) {
	var (
		i, j int
	)

	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[i]); j++ {
			if grid[i][j] == bempty {
				bombDetSecs[bdetKey(i, j)] = bdetsecs
				grid[i][j] = bbomb
			}
		}
	}
}

func printGrid(grid [][]rune) {
	for _, v := range grid {
		for _, char := range v {
			fmt.Printf("%c", char)
		}

		fmt.Println()
	}

	fmt.Println()
}

func convertGrid(grid [][]rune) []string {
	newGrid := make([]string, len(grid))

	for i, row := range grid {
		newGrid[i] = string(row)
	}

	return newGrid
}

// Bomberman FIXME
func Bomberman(n int, gridIn []string) []string {
	var (
		i, j int
		grid [][]rune
	)

	// create the rune grid
	grid = make([][]rune, len(gridIn))
	for i = range gridIn {
		grid[i] = make([]rune, len(gridIn[i]))
		j = 0
		for _, char := range gridIn[i] {
			grid[i][j] = char
			j++
		}
	}

	// printGrid(grid)

	bombDetSecs := map[string]int{}

	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[i]); j++ {
			if grid[i][j] == bbomb {
				bombDetSecs[bdetKey(i, j)] = bdetsecs
			}
		}
	}

	// Step 2
	detBombs(bombDetSecs, grid)
	n--
	// printGrid(grid)

	for n > 0 {
		// Step 3
		detBombs(bombDetSecs, grid)
		placeBombs(bombDetSecs, grid)
		n--
		// printGrid(grid)

		// Step 4
		if n > 0 {
			detBombs(bombDetSecs, grid)
			n--
			// printGrid(grid)
		}
	}

	return convertGrid(grid)
}
