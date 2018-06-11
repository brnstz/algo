package algo

// Graph FIXME
type Graph struct {
	Edges []*Edge

	// Adj maps vertices to a slice of all the other vertices they
	// are adjacent to.
	Adj map[*Vertex][]*Vertex
}

// Vertex FIXME
type Vertex struct {
	Value interface{}
}

// Edge FIXME
type Edge struct {
	Weight float64

	From, To *Vertex
}

func (edge *Edge) PQLess(other PQItem) {
	otherEdge := (*Edge).other
	return e.Weight > otherEdge.Weight
}

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
func (g *Graph) MinimumSpanningTree(v *Vertex) []*Edge {
	var (
		mst        []*Edge
		edge       *Edge
		goodEdge   *Edge
		edgePQ     *PriorityQueue
		nextEdgePQ *PriorityQueue
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
			return nil
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
			edge = (*Edge).edgePQ.DelMax()

			if goodEdge == nil && marked[edge.From] != marked[edge.To] {
				// Find the lowest weight edge that expands our tree. It's
				// something that is on our tree *and* expands it. That is, one
				// vertex is marked but the other isn't.
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

	return mst
}
