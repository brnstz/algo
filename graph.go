package algo

import (
	"fmt"
	"math"
)

// Graph is a system of Vertices connected via Edges
type Graph struct {
	Edges []*Edge

	// Adj maps vertices to a slice of all the other edges connected to it.
	Adj map[*Vertex][]*Edge
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

// Path defines a way to get from Edges[0].From to Edges[len(Edges)-1].To
// along with the Weight (or cost / distance, etc.) of going there.
type Path struct {
	Weight   float64
	From, To *Vertex
	Edges    []*Edge
}

func (p *Path) String() string {
	if len(p.Edges) < 1 {
		return "No Path"
	}

	x := fmt.Sprintf("From %v to %v, weight: %v, Path: [\n",
		p.From, p.To, p.Weight,
	)

	for _, edge := range p.Edges {
		x += edge.String() + "\n"
	}

	x += "]"

	return x
}

// AddEdge creates a connection on the Graph from edge.From to edge.To
func (g *Graph) AddEdge(edge *Edge) {
	// Initialize adjacency mapping if necessary
	if g.Adj == nil {
		g.Adj = map[*Vertex][]*Edge{}
	}

	// Ensure that any vertex with no edge out exists in the map
	if g.Adj[edge.To] == nil {
		g.Adj[edge.To] = make([]*Edge, 0)
	}

	// Record that this vertex has another edge
	g.Adj[edge.From] = append(g.Adj[edge.From], edge)

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

type vertexWeight struct {
	vertex *Vertex
	weight float64
}

// PQLess implements the PQItem to allow us to use edges on a PriorityQueue
func (vw *vertexWeight) PQLess(other PQItem) bool {
	otherVW := other.(*vertexWeight)
	return vw.weight > otherVW.weight
}

func (vw *vertexWeight) String() string {
	return fmt.Sprintf("%v", vw.weight)
}

// ShortestPath returns a mapping of the shortest path from source to every
// connected vertex in the Graph
func (g *Graph) ShortestPath(source *Vertex) (map[*Vertex]*Path, error) {
	var (
		err    error
		weight float64
		i      int

		item   PQItem
		vw     *vertexWeight
		vertex *Vertex
		next   *Vertex
		edge   *Edge
	)

	visited := map[*Vertex]bool{}
	edgeTo := map[*Vertex]*Edge{}
	weightTo := map[*Vertex]*vertexWeight{}
	paths := map[*Vertex]*Path{}

	vwPQ := NewPriorityQueue(len(g.Adj))

	for vertex = range g.Adj {

		if vertex == source {
			// If this is the same vertex, weight is 0
			weight = 0.0
		} else {
			// Otherwise it's infinite
			weight = math.MaxFloat64
		}

		// Create a vertexWeight to put onto the priorityQueue
		vw := &vertexWeight{
			vertex: vertex,
			weight: weight,
		}

		weightTo[vertex] = vw

		// Add this vertexWeight to our queue
		vwPQ.Insert(vw)
	}

	// Process vertices from lowest to highest weight
	for !vwPQ.IsEmpty() {

		// Coerce the pq item into a vertexWeight
		item, err = vwPQ.DelMax()
		if err != nil {
			return nil, err
		}
		vw = item.(*vertexWeight)

		// Ignore vertices we've already visited
		if visited[vw.vertex] {
			continue
		}

		// Mark that we have visited this vertex
		visited[vw.vertex] = true

		for _, edge = range g.Adj[vw.vertex] {

			// Ignore vertices we've already visited
			if visited[edge.To] {
				continue
			}

			// What is the weight if we use this edge?
			weight = weightTo[vw.vertex].weight + edge.Weight

			if weight < weightTo[edge.To].weight {

				// If that weight is less that the current weight to
				// edge.To, then use it.
				weightTo[edge.To].weight = weight
				edgeTo[edge.To] = edge

				// Register the weight change with our priority queue
				i, err = vwPQ.IndexOf(weightTo[edge.To])
				if err != nil {
					return nil, err
				}
				vwPQ.IndicateChange(i)

			}
		}
	}

	// Create a Path object for each vertex that isn't the source
	for vertex = range g.Adj {
		if vertex == source {
			continue
		}

		// Initialize path with the weight and the source/destination
		path := &Path{
			Weight: weightTo[vertex].weight,
			From:   source,
			To:     vertex,
		}

		// Initialize next vertex to the destination
		next = vertex
		for {

			// Get the next edge we need
			edge = edgeTo[next]

			// Append to our list of edges
			path.Edges = append(path.Edges, edge)

			// If we made it to the source, we are done
			if edge.From == source {
				break
			}

			// Otherwise, go to the From side of this edge
			next = edge.From
		}

		// We've found the path, map it to the destination
		paths[vertex] = path
	}

	return paths, nil
}
