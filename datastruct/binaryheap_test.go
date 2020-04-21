package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryHeap(t *testing.T) {
	in := []int{7, 3, 8, 4, 1, 5, 5}
	expected := []int{1, 3, 4, 5, 5, 7, 8}
	bh := &BinaryHeap{[]int{}}
	for _, v := range in {
		bh.Insert(v)
	}
	assert.Equal(t, 1, bh.Min())
	assert.Equal(t, expected, bh.Heapsort([]int{}))
}
