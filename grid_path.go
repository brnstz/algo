package algo

import (
	"fmt"
	"log"
)

func findVertex(vertices [][]*Vertex, x, y int) *Vertex {

	if x >= len(vertices) || x < 0 {
		// Bad x coord
		return nil

	} else if y >= len(vertices[x]) || y < 0 {
		// Bad x, y coord
		return nil

	} else if vertices[x][y] == nil {
		// The endpoint is a wall
		return nil
	}

	// It's good
	return vertices[x][y]
}

func createVertices(grid [][]bool) [][]*Vertex {
	var (
		i, j int
	)

	vertices := make([][]*Vertex, len(grid))

	for i = 0; i < len(grid); i++ {
		vertices[i] = make([]*Vertex, len(grid[i]))

		for j = 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				vertices[i][j] = &Vertex{
					Value: fmt.Sprintf("(%v,%v)", i, j),
				}
			}
		}
	}

	return vertices
}

// GridPath returns the shortest path from (sx, sy) to (ex, ey) within
// grid. A true entry in the grid is a valid path, where as false is
// not (e.g., it's a wall). All paths have equal weight and all paths
// are two dimensional.
func GridPath(grid [][]bool, sx, sy, ex, ey int) *Path {
	var (
		i, j, x, y int
		d          []int
		edge       *Edge
		start, end *Vertex
		from, to   *Vertex
	)

	g := Graph{}
	vertices := createVertices(grid)

	// Given an x, y position, these diffs to x, y are all possible
	// adjacent points.
	diffs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	start = findVertex(vertices, sx, sy)
	end = findVertex(vertices, ex, ey)

	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[i]); j++ {

			from = findVertex(vertices, i, j)
			if from == nil {
				continue
			}

			for _, d = range diffs {
				x = i + d[0]
				y = j + d[1]

				to = findVertex(vertices, x, y)
				if to == nil {
					continue
				}

				edge = &Edge{
					From:   from,
					To:     to,
					Weight: 1,
				}
				g.AddEdge(edge)

			}
		}
	}

	paths, err := g.ShortestPath(start)
	if err != nil {
		log.Printf("error finding path from %v to %v: %v", start, end, err)
		return nil
	}

	return paths[end]
}
