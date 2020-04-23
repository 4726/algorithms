package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLTreeLeftRight1(t *testing.T) {
	avl := &AVLTree{}
	avl.Insert(20)
	avl.Insert(4)
	avl.Insert(15)
	assert.Equal(t, 15, avl.root.value)
	assert.Equal(t, 4, avl.root.left.value)
	assert.Equal(t, 20, avl.root.right.value)
}

func TestAVLTreeLeftRight2(t *testing.T) {
	avl := &AVLTree{}
	avl.Insert(20)
	avl.Insert(4)
	avl.Insert(26)
	avl.Insert(3)
	avl.Insert(9)
	avl.Insert(15)
	assert.Equal(t, 9, avl.root.value)
	assert.Equal(t, 4, avl.root.left.value)
	assert.Equal(t, 20, avl.root.right.value)
}

func TestAVLTreeLeftRight3(t *testing.T) {
	avl := &AVLTree{}
	avl.Insert(20)
	avl.Insert(4)
	avl.Insert(26)
	avl.Insert(3)
	avl.Insert(9)
	avl.Insert(21)
	avl.Insert(30)
	avl.Insert(2)
	avl.Insert(7)
	avl.Insert(11)
	avl.Insert(15)
	assert.Equal(t, 9, avl.root.value)
	assert.Equal(t, 4, avl.root.left.value)
	assert.Equal(t, 20, avl.root.right.value)
}

func TestAVLTreeLeftLeft(t *testing.T) {
	avl := &AVLTree{}
	avl.Insert(13)
	avl.Insert(10)
	avl.Insert(15)
	avl.Insert(5)
	avl.Insert(11)
	avl.Insert(16)
	avl.Insert(4)
	avl.Insert(8)
	avl.Insert(3)
	assert.Equal(t, 13, avl.root.value)
	assert.Equal(t, 5, avl.root.left.value)
	assert.Equal(t, 15, avl.root.right.value)
}

func TestAVLTreeRightRight(t *testing.T) {
	avl := &AVLTree{}
	avl.Insert(30)
	avl.Insert(5)
	avl.Insert(35)
	avl.Insert(32)
	avl.Insert(40)
	avl.Insert(45)
	assert.Equal(t, 35, avl.root.value)
	assert.Equal(t, 30, avl.root.left.value)
	assert.Equal(t, 40, avl.root.right.value)
}

func TestAVLTreeRightLeft(t *testing.T) {
	avl := &AVLTree{}
	avl.Insert(5)
	avl.Insert(2)
	avl.Insert(7)
	avl.Insert(1)
	avl.Insert(4)
	avl.Insert(6)
	avl.Insert(9)
	avl.Insert(3)
	avl.Insert(16)
	avl.Insert(15)
	assert.Equal(t, 5, avl.root.value)
	assert.Equal(t, 2, avl.root.left.value)
	assert.Equal(t, 7, avl.root.right.value)
}
