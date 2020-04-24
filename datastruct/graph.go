package datastruct

//vertex = node
//edge = line connecting the nodes

type Graph struct {
	vertices map[int]map[int]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		vertices: map[int]map[int]struct{}{},
	}
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

func (g *Graph) BreadthFirstSearch(root int) []int {
	var res []int
	visited := map[int]struct{}{}

	queue := []int{root}
	for len(queue) > 0 {
		if _, ok := visited[queue[0]]; ok {
			if len(queue) == 1 {
				queue = []int{}
			} else {
				queue = queue[1:]
			}
			continue
		}
		res = append(res, queue[0])
		visited[queue[0]] = struct{}{}
		edges, ok := g.vertices[queue[0]]
		if !ok {
			return res
		}

		if len(queue) == 1 {
			queue = []int{}
		} else {
			queue = queue[1:]
		}

		for vertex := range edges {
			if _, ok := visited[vertex]; !ok {
				queue = append(queue, vertex)
			}
		}
	}
	return res
}

func (g *Graph) DepthFirstSearch(root int) []int {
	var res []int
	visited := map[int]struct{}{}

	stack := []int{root}
	for len(stack) > 0 {
		index := len(stack) - 1
		if _, ok := visited[stack[index]]; ok {
			if len(stack) == 1 {
				stack = []int{}
			} else {
				stack = stack[:index]
			}
			continue
		}
		res = append(res, stack[index])
		visited[stack[index]] = struct{}{}
		edges, ok := g.vertices[stack[index]]
		if !ok {
			return res
		}

		if len(stack) == 1 {
			stack = []int{}
		} else {
			stack = stack[:index]
		}

		for vertex := range edges {
			if _, ok := visited[vertex]; !ok {
				stack = append(stack, vertex)
			}
		}
	}
	return res
}
