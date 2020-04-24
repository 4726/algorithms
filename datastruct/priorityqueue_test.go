package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	in := []int{7, 3, 8, 4, 1, 5, 5}
	pq := NewPriorityQueue()
	for _, v := range in {
		pq.Insert(v)
	}
	assert.Equal(t, 8, pq.Pop())
	assert.Equal(t, 7, pq.Pop())
	assert.Equal(t, 5, pq.Pop())
	assert.Equal(t, 5, pq.Pop())
	assert.Equal(t, 4, pq.Pop())
	assert.Equal(t, 3, pq.Pop())
	assert.Equal(t, 1, pq.Pop())
}
