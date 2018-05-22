package algo_test

import (
	"github.com/brnstz/algo"

	"fmt"
	"io"
	"math"
	"os"
	"testing"
)

// Implement NodeValue interface for strings
type stringNode string

func (s stringNode) Less(other_ algo.NodeValue) bool {
	other := other_.(stringNode)
	return s < other
}

func (s stringNode) Equals(other_ algo.NodeValue) bool {
	other := other_.(stringNode)
	return s == other
}

func TestRedBlack(t *testing.T) {
	tree := algo.RedBlackTree{}

	fh, err := os.Open("data/tale.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()

	var word stringNode
	for {
		_, err := fmt.Fscan(fh, &word)
		if err == io.EOF {
			break
		}

		tree.Put(word)
	}

	// Tree should have some words but not others
	var yesFind, noFind stringNode
	yesFind = "goodfellowship"
	noFind = "slfkjkldsf"

	if tree.Find(yesFind) != true {
		t.Fatal("Cannot find word")
	}

	if tree.Find(noFind) != false {
		t.Fatal("Found unexpected word")
	}

	height := tree.Height()

	// A red black tree should have at most 2log(n + 1) height
	maxNodes := 2 * math.Log2(float64(tree.Root.NodeCount+1))

	if float64(height) > maxNodes {
		t.Fatalf("Tree is too high, actual: %v, expected < %v", height, maxNodes)
	}

	out := tree.BFSString()
	fmt.Print(out)
	fmt.Println("Tree height: ", tree.Height())
	fmt.Println("Node count: ", tree.Root.NodeCount)

}
