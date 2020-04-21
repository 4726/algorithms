package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	in := []int{7, 3, 8, 4, 1, 5, 5}
	expected := []int{1, 3, 4, 5, 5, 7, 8}
	expectedRev := []int{8, 7, 5, 5, 4, 3, 1}
	bst := &BinarySearchTree{nil}
	for _, v := range in {
		bst.Insert(v)
	}
	assert.Equal(t, expected, bst.Sorted())
	assert.Equal(t, expectedRev, bst.SortedRev())
}
