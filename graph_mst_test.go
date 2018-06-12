package algo

import (
	"fmt"
	"testing"
)

func TestMST(t *testing.T) {
	var (
		i, j, newI, newJ int
		result           []*Edge
		err              error
		totalWeight      float64
	)

	g := Graph{}
	maze := [][]int{
		{1, 100, 200, 400},
		{10, 50, 2, 3},
		{1, 1, 1, 1},
	}

	vertices := make([][]*Vertex, len(maze))

	// Create a Vertex object for each element of the maze
	for i = 0; i < len(maze); i++ {
		vertices[i] = make([]*Vertex, len(maze[i]))
		for j = 0; j < len(maze[i]); j++ {
			vertices[i][j] = &Vertex{Value: fmt.Sprintf("(%v,%v)", i, j)}
		}
	}

	// Connect adjacent parts of the maze by using the value of maze[i][j]
	// as the "weight" of going From some vetext To vertex[i][j]. That is,
	// the vertex you are going *to* contains the weight of travel.
	for i = 0; i < len(maze); i++ {
		for j = 0; j < len(maze[i]); j++ {

			// Try every possible adjacent vertex.
			for _, diff := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				newI = i + diff[0]
				newJ = j + diff[1]

				// If it's valid vertex, add any edge.
				if newI >= 0 && newI < len(maze) &&
					newJ >= 0 && newJ < len(maze[i]) {
					g.AddEdge(
						&Edge{
							From:   vertices[i][j],
							To:     vertices[newI][newJ],
							Weight: float64(maze[newI][newJ]),
						},
					)
				}
			}
		}
	}

	result, err = g.MinimumSpanningTree(vertices[0][0])
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, edge := range result {
		totalWeight += edge.Weight
	}

	if totalWeight != 14 {
		t.Fatalf("expected totalWeight == 14 but got %v", totalWeight)
	}
}
