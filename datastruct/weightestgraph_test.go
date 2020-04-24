package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	graph := NewWeightedGraph()
	for i := 0; i < 6; i++ {
		graph.AddVertex(i)
	}
	graph.AddEdge(0, 1, 4)
	graph.AddEdge(0, 2, 4)
	graph.AddEdge(1, 2, 2)
	graph.AddEdge(2, 3, 3)
	graph.AddEdge(2, 4, 6)
	graph.AddEdge(2, 5, 1)
	graph.AddEdge(3, 4, 2)
	graph.AddEdge(5, 4, 3)
	distance := graph.Dijkstra(0, 4)
	assert.Equal(t, 8, distance)
}
