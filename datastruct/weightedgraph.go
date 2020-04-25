package datastruct

import "math"

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

func (g *WeightedGraph) AddEdge(x, y, weight int) {
	if _, ok := g.vertices[x]; !ok {
		return
	}
	if _, ok := g.vertices[y]; !ok {
		return
	}

	g.vertices[x][y] = weight
	g.vertices[y][x] = weight
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
	previous := map[int]int{}
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
				previous[key] = leastKey
			}
		}
	}

	var routeRev []int
	vertexKey := destination
	for vertexKey != source {
		routeRev = append(routeRev, vertexKey)
		prevKey := previous[vertexKey]
		vertexKey = prevKey
	}
	routeRev = append(routeRev, source)
	var route []int
	for i := len(routeRev) - 1; i > -1; i-- {
		route = append(route, routeRev[i])
	}

	return distances[destination], route
}
