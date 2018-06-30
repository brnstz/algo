package puzzles

import (
	"math"
)

// https://www.hackerrank.com/challenges/cut-the-tree/problem

type tree struct {
	edges          [][]int
	adj            [][]int
	weights        []int
	visited        []bool
	subtreeWeights []int
}

func assignSubtreeWeight(tree *tree, node int) int {
	tree.visited[node] = true

	tree.subtreeWeights[node] = tree.weights[node]

	for _, child := range tree.adj[node] {
		if tree.visited[child] {
			continue
		}

		tree.subtreeWeights[node] += assignSubtreeWeight(tree, child)
	}

	return tree.subtreeWeights[node]
}

// CutTree finds the minimal difference when cutting the tree in two halves
func CutTree(weights []int, edges [][]int) int {
	tree := &tree{
		edges:          edges,
		weights:        weights,
		subtreeWeights: make([]int, len(weights)),
		visited:        make([]bool, len(weights)),
		adj:            make([][]int, len(weights)),
	}

	for _, edge := range tree.edges {
		tree.adj[edge[0]-1] = append(tree.adj[edge[0]-1], edge[1]-1)
		tree.adj[edge[1]-1] = append(tree.adj[edge[1]-1], edge[0]-1)
	}

	assignSubtreeWeight(tree, 0)

	minDiff := math.MaxInt32

	for i := 1; i < len(tree.subtreeWeights); i++ {
		diff := int(math.Abs(float64(tree.subtreeWeights[0] - (tree.subtreeWeights[i] * 2))))
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
