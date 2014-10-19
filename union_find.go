package algo

type UnionFind struct {
	// Number of nodes in the system. Static on init.
	nodeCount int

	// Number of disconnected trees, decreases as nodes become connected.
	treeCount int

	// Mapping from the node to its parent. Each node starts with a
	// link to itself, but this changes as nodes are connected. If a
	// parent[i] == i, then i is a currently root node.
	parent []int

	// For nodes that are roots, the size (number of nodes) in their tree.
	// Orphan values get left here when a former root nodes becomes part
	// of a larger tree.
	size []int
}

// Create a new union find structure which can hold count components
func NewUnionFind(count int) *UnionFind {
	uf := &UnionFind{}

	uf.nodeCount = count
	uf.treeCount = count

	uf.parent = make([]int, count)
	uf.size = make([]int, count)

	for i := 0; i < count; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}

	return uf
}

// Find the tree which node is in
func (uf *UnionFind) Find(node int) int {
	// Keep finding until we get to a self-connected component (a root node)
	for node != uf.parent[node] {
		node = uf.parent[node]
	}

	return node
}

// Connect node1 and node2
func (uf *UnionFind) Union(node1, node2 int) {

	// Find the root of each node
	root1 := uf.Find(node1)
	root2 := uf.Find(node2)

	// If same root, then nodes are already connected.
	if root1 == root2 {
		return
	}

	// Find the larger tree, and connect smaller tree to it.
	if uf.size[root1] > uf.size[root2] {
		// root1 is bigger, connect root2 to it.
		uf.parent[root2] = root1

		// The size of the tree at root1 increases by the existing size of
		// root2's tree. uf.size[root2] is now a meaningless orphan value.
		uf.size[root1] += uf.size[root2]
	} else {
		// Same as last clause but reverse
		uf.parent[root1] = root2
		uf.size[root2] += uf.size[root1]
	}

	// Since we combined two trees, there is now one less
	uf.treeCount--
}

// Are these two nodes in the same tree?
func (uf *UnionFind) Connected(node1, node2 int) bool {
	return uf.Find(node1) == uf.Find(node2)
}

func (uf *UnionFind) Parent() []int {
	return uf.parent
}
