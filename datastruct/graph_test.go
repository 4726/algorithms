package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreadthFirstSearch(t *testing.T) {
	graph := NewGraph()
	for i := 0; i < 9; i++ {
		graph.AddVertex(i)
	}
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)
	graph.AddEdge(0, 8)
	graph.AddEdge(1, 7)
	graph.AddEdge(8, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(3, 2)
	graph.AddEdge(2, 5)
	graph.AddEdge(5, 6)
	graph.AddEdge(2, 7)
	bfs := graph.BreadthFirstSearch(0)

	assert.Equal(t, 0, bfs[0])
	assert.ElementsMatch(t, []int{1, 3, 8}, bfs[1:4])
	assert.ElementsMatch(t, []int{7, 2, 4}, bfs[4:7])
	assert.Equal(t, 5, bfs[7])
	assert.Equal(t, 6, bfs[8])
}

func TestDepthFirstSearch(t *testing.T) {
	graph := NewGraph()
	for i := 0; i < 9; i++ {
		graph.AddVertex(i)
	}
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)
	graph.AddEdge(0, 8)
	graph.AddEdge(1, 7)
	graph.AddEdge(8, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(3, 2)
	graph.AddEdge(2, 5)
	graph.AddEdge(5, 6)
	graph.AddEdge(2, 7)
	dfs := graph.DepthFirstSearch(0)

	expectedAny := [][]int{
		[]int{0, 8, 4, 3, 2, 5, 6, 7, 1},
		[]int{0, 8, 4, 3, 2, 7, 1, 5, 6},
		[]int{0, 3, 4, 8, 2, 7, 1, 5, 6},
		[]int{0, 3, 4, 8, 2, 5, 6, 7, 1},
		[]int{0, 3, 2, 7, 1, 5, 6, 4, 8},
		[]int{0, 3, 2, 5, 6, 7, 1, 4, 8},
		[]int{0, 1, 7, 2, 3, 4, 8, 5, 6},
		[]int{0, 1, 7, 2, 5, 6, 3, 4, 8},
	}

	assert.Contains(t, expectedAny, dfs)
}
