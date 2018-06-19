package algo

import (
	"testing"
)

func TestShortestPath(t *testing.T) {
	var err error

	g := Graph{}
	v1 := &Vertex{Value: "Boston"}
	v2 := &Vertex{Value: "New York"}
	v3 := &Vertex{Value: "Philadelphia"}
	v4 := &Vertex{Value: "Baltimore"}
	v5 := &Vertex{Value: "Washington, DC"}

	g.AddEdge(
		&Edge{
			From: v1, To: v2, Weight: 230,
		},
	)
	g.AddEdge(
		&Edge{
			From: v2, To: v3, Weight: 99,
		},
	)
	g.AddEdge(
		&Edge{
			From: v3, To: v4, Weight: 105,
		},
	)
	g.AddEdge(
		&Edge{
			From: v4, To: v5, Weight: 40,
		},
	)
	g.AddEdge(
		&Edge{
			From: v1, To: v5, Weight: 101,
		},
	)
	g.AddEdge(
		&Edge{
			From: v5, To: v3, Weight: 10,
		},
	)
	g.AddEdge(
		&Edge{
			From: v5, To: v3, Weight: 5,
		},
	)

	g.AddEdge(
		&Edge{
			From: v1, To: v2, Weight: 3,
		},
	)

	paths, err := g.ShortestPath(v1)
	if err != nil {
		t.Fatal(err)
	}

	if paths[v2].Weight != 3 {
		t.Fatalf("Expected path of weight 3 but got this path: %v", paths[v2])
	}

	if paths[v3].Weight != 102 {
		t.Fatalf("Expected path of weight 102 but got this path: %v", paths[v3])
	}

}
