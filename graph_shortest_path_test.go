package algo

import "testing"

func TestShortestPath(t *testing.T) {

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

	g.ShortestPath(v1)
	t.Fatal()

}
