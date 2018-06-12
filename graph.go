package algo

import (
	"fmt"
)

// Graph is a system of Vertices connected via Edges
type Graph struct {
	Edges []*Edge

	// Adj maps vertices to a slice of all the other vertices they
	// are adjacent to.
	Adj map[*Vertex][]*Vertex
}

// Vertex is a part of a Graph that is connected to other vertices via
// Edges.
type Vertex struct {
	Value interface{}
}

func (v *Vertex) String() string {
	return fmt.Sprintf("%v", v.Value)
}

// Edge is a connection between two Verticies with a particular Weight.
// Connections are one directional. To create uni-directional graphs, use
// one Edge for each direction.
type Edge struct {
	Weight float64

	From, To *Vertex
}

func (e *Edge) String() string {
	return fmt.Sprintf("%v => %v @ %v", e.From.Value, e.To.Value, e.Weight)
}

// PQLess implements the PQItem to allow us to use edges on a PriorityQueue
func (e *Edge) PQLess(other PQItem) bool {
	otherEdge := other.(*Edge)
	return e.Weight > otherEdge.Weight
}

// AddEdge creates a connection on the Graph from edge.From to edge.To
func (g *Graph) AddEdge(edge *Edge) {
	// Initialize adjacency mapping if necessary
	if g.Adj == nil {
		g.Adj = map[*Vertex][]*Vertex{}
	}

	// Record that the from vertex is adjacent to the To vertex
	g.Adj[edge.From] = append(g.Adj[edge.From], edge.To)

	// Add this edge to our slice of edges
	g.Edges = append(g.Edges, edge)
}

// MinimumSpanningTree finds the minimal list of edges to span the entire
// graph that is connected to v. If v == nil, we use the first Edge.From
// value.
func (g *Graph) MinimumSpanningTree(v *Vertex) ([]*Edge, error) {
	var (
		mst        []*Edge
		edge       *Edge
		goodEdge   *Edge
		edgePQ     *PriorityQueue
		nextEdgePQ *PriorityQueue
		err        error
		pqItem     PQItem
	)

	// Mark which vertices we have visited
	marked := map[*Vertex]bool{}

	// Initialize edgePQ with all edges
	edgePQ = NewPriorityQueue(len(g.Edges))
	for _, edge := range g.Edges {
		edgePQ.Insert(edge)
	}

	// If no vertex passed in, then assume the first one
	if v == nil {
		if len(g.Edges) > 0 {

			// If we have edges, then use the first From value
			v = g.Edges[0].From
		} else {

			// Otherwise there is no solution
			return nil, nil
		}
	}

	// Mark the first vertex as visited
	marked[v] = true

	for !edgePQ.IsEmpty() {
		// Create a new PQ for the next Loop
		nextEdgePQ = NewPriorityQueue(len(g.Edges))

		// Find one goodEdge per loop below
		goodEdge = nil

		for !edgePQ.IsEmpty() {

			// Get the lowest weight edge on the priority queue
			pqItem, err = edgePQ.DelMax()
			edge = pqItem.(*Edge)
			if err != nil {
				return nil, err
			}

			if goodEdge == nil && marked[edge.From] != marked[edge.To] {
				// Find the lowest weight edge that expands our tree.  That is,
				// one vertex is marked but the other isn't.
				goodEdge = edge

			} else {

				// Otherwise, add it to the next queue
				nextEdgePQ.Insert(edge)
			}
		}

		// If we didn't find a goodEdge, then we're done
		if goodEdge == nil {
			break
		}

		// Save that each vertex for this Edge is marked
		marked[goodEdge.From] = true
		marked[goodEdge.To] = true

		// Add it to the list of edges for the solution
		mst = append(mst, goodEdge)

		// Use the new queue for next round
		edgePQ = nextEdgePQ
	}

	return mst, nil
}
