package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarysearch(t *testing.T) {
	in := []int{1, 3, 4, 5, 5, 7, 8}
	index := Binarysearch(in, 8)
	assert.Equal(t, 6, index)
	index = Binarysearch(in, 1)
	assert.Equal(t, 0, index)
	index = Binarysearch(in, 6)
	assert.Equal(t, -1, index)
}
