package algo

// https://www.hackerrank.com/challenges/queens-attack-2/problem

const (
	empty = iota
	queen
	obstacle
	// path
)

func validMove(x, y int, grid [][]int) bool {
	return x >= 0 && y >= 0 &&
		x < len(grid) && y < len(grid[x]) &&
		grid[x][y] == empty
}

// QueenAttack FIXME
func QueenAttack(n, numObstacles, x, y int, obstacles [][]int) int {
	var (
		squares   int
		ax, ay, i int
		grid      [][]int
	)

	// Initialize grid with empty (0) values
	grid = make([][]int, n)
	for i = range grid {
		grid[i] = make([]int, n)
	}

	// Set position of the queen
	grid[x-1][y-1] = queen

	// Set position of obstacles
	for _, v := range obstacles {
		grid[v[0]-1][v[1]-1] = obstacle
	}

	attackDirs := [][]int{
		{0, 1}, {0, -1},
		{1, 0}, {-1, 0},
		{1, 1}, {-1, -1},
		{-1, 1}, {1, -1},
	}

	squares = 0
	for _, attack := range attackDirs {
		ax = x - 1
		ay = y - 1

		for true {
			ax += attack[0]
			ay += attack[1]

			if validMove(ax, ay, grid) {
				squares++
			} else {
				break
			}
		}
	}

	return squares
}
