package puzzles

import (
	"fmt"
	"log"
	"math"
)

// https://www.hackerrank.com/challenges/cut-the-tree/problem

type treeNode struct {
	weight        int
	subtreeWeight int
	children      []*treeNode
	connected     bool
}

func (n *treeNode) String() string {
	return fmt.Sprintf(
		"weight: %v, subtreeWeight: %v, children: %v",
		n.weight, n.subtreeWeight, len(n.children),
	)
}

func assignSubtreeWeight(node *treeNode) int {
	var (
		child *treeNode
	)

	node.subtreeWeight = node.weight
	for _, child = range node.children {
		node.subtreeWeight += assignSubtreeWeight(child)
	}

	return node.subtreeWeight
}

// CutTree FIXME
func CutTree(weights []int, edges [][]int) int {
	var (
		i, weight, diff, minDiff int
		node                     *treeNode
		node1, node2             *treeNode
		nodes                    []*treeNode
		edge                     []int
	)

	nodes = make([]*treeNode, len(weights))

	for i, weight = range weights {
		nodes[i] = &treeNode{weight: weight}
	}

	for _, edge = range edges {
		node1 = nodes[edge[0]-1]
		node2 = nodes[edge[1]-1]

		if node1.connected {
			node1.children = append(node1.children, node2)

		} else {
			log.Printf("what about node2? %v %v", node2.connected, node2)
			node2.children = append(node2.children, node1)
		}

		node1.connected = true
		node2.connected = true
	}

	assignSubtreeWeight(nodes[0])

	minDiff = math.MaxInt32

	for _, node = range nodes {
		diff = int(math.Abs(float64(nodes[0].subtreeWeight - (node.subtreeWeight * 2))))
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
