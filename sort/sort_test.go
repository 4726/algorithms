package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	in := []int{7, 3, 8, 4, 1, 5, 5}
	expected := []int{1, 3, 4, 5, 5, 7, 8}
	Quicksort(in)
	assert.Equal(t, expected, in)
}

func TestMergesort(t *testing.T) {
	in := []int{7, 3, 8, 4, 1, 5, 5}
	expected := []int{1, 3, 4, 5, 5, 7, 8}
	Mergesort(in)
	assert.Equal(t, expected, in)
}

func TestInsertionsort(t *testing.T) {
	in := []int{7, 3, 8, 4, 1, 5, 5}
	expected := []int{1, 3, 4, 5, 5, 7, 8}
	Insertionsort(in)
	assert.Equal(t, expected, in)
}

func TestSelectionsort(t *testing.T) {
	in := []int{7, 3, 8, 4, 1, 5, 5}
	expected := []int{1, 3, 4, 5, 5, 7, 8}
	Selectionsort(in)
	assert.Equal(t, expected, in)
}

func TestBubblesort(t *testing.T) {
	in := []int{7, 3, 8, 4, 1, 5, 5}
	expected := []int{1, 3, 4, 5, 5, 7, 8}
	Bubblesort(in)
	assert.Equal(t, expected, in)
}
