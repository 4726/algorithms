package datastruct

//vertex = node
//edge = line connecting the nodes

type Graph struct {
	vertices map[int]map[int]struct{}
}

func (g *Graph) Adjacent(x, y int) bool {
	edges := g.vertices[x]
	_, ok := edges[y]
	return ok
}

func (g *Graph) Neighbors(x int) map[int]struct{} {
	return g.vertices[x]
}

func (g *Graph) AddVertex(x int) {
	if _, ok := g.vertices[x]; !ok {
		g.vertices[x] = map[int]struct{}{}
	}
}

func (g *Graph) RemoveVertex(x int) {
	delete(g.vertices, x)
	for _, edges := range g.vertices {
		delete(edges, x)
	}
}

func (g *Graph) AddEdge(x, y int) {
	if _, ok := g.vertices[x]; !ok {
		return
	}
	if _, ok := g.vertices[y]; !ok {
		return
	}

	g.vertices[x][y] = struct{}{}
	g.vertices[y][x] = struct{}{}
}

func (g *Graph) RemoveEdge(x, y int) {
	edges, ok := g.vertices[x]
	if !ok {
		return
	}
	if _, ok = g.vertices[y]; !ok {
		return
	}

	delete(edges, y)
	g.vertices[x] = edges
	g.RemoveEdge(y, x)
}
