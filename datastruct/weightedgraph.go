package datastruct

import "math"

//vertex = node
//edge = line connecting the nodes

type WeightedGraph struct {
	vertices map[int]map[int]int
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{
		vertices: map[int]map[int]int{},
	}
}

func (g *WeightedGraph) AddVertex(x int) {
	if _, ok := g.vertices[x]; !ok {
		g.vertices[x] = map[int]int{}
	}
}

func (g *WeightedGraph) RemoveVertex(x int) {
	delete(g.vertices, x)
	for _, edges := range g.vertices {
		delete(edges, x)
	}
}

func (g *WeightedGraph) AddEdge(x, y, len int) {
	if _, ok := g.vertices[x]; !ok {
		return
	}
	if _, ok := g.vertices[y]; !ok {
		return
	}

	g.vertices[x][y] = len
	g.vertices[y][x] = len
}

func (g *WeightedGraph) RemoveEdge(x, y int) {
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

func (g *WeightedGraph) Dijkstra(source, destination int) (int, []int) {
	distances := map[int]int{}
	vertexKeys := map[int]struct{}{}
	routes := map[int]int{}
	for vertexKey := range g.vertices {
		distances[vertexKey] = math.MaxInt64
		vertexKeys[vertexKey] = struct{}{}
	}
	distances[source] = 0

	for len(vertexKeys) > 0 {
		least := math.MaxInt64
		var leastKey int
		for vertexKey := range vertexKeys {
			if dist := distances[vertexKey]; dist < least {
				leastKey = vertexKey
				least = dist
			}
		}
		delete(vertexKeys, leastKey)

		edges := g.vertices[leastKey]
		for key, distance := range edges {
			totalDistance := distances[leastKey] + distance
			if totalDistance < distances[key] {
				distances[key] = totalDistance
				routes[leastKey] = key
			}
		}
	}

	route := []int{source}
	for from, to := range routes {
		if route[len(route)-1] == from {
			route = append(route, to)
		} else {
			route = append(route, from)
		}
	}
	return distances[destination], route
}
