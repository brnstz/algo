package algo_test

import (
	"github.com/brnstz/algo"

	"fmt"
	"io"
	"os"
	"testing"
)

func TestUnionFind(t *testing.T) {
	fh, err := os.Open("data/mediumUF.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()

	var numNodes, node1, node2 int
	fmt.Fscan(fh, &numNodes)

	uf := algo.NewUnionFind(numNodes)

	for {
		_, err := fmt.Fscan(fh, &node1, &node2)
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}

		uf.Union(node1, node2)
	}

	if !uf.Connected(403, 452) {
		t.Fatal("Expected nodes not connected")
	}

	if uf.Connected(100, 305) {
		t.Fatal("Nodes are unexpectedly connected")
	}
}
