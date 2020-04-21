package datastruct

type BinarySearchTree struct {
	root *BSTNode
}

type BSTNode struct {
	left, right *BSTNode
	value       int
}

func (bst *BinarySearchTree) Insert(v int) {
	if bst.root == nil {
		bst.root = &BSTNode{nil, nil, v}
		return
	}

	newNode := &BSTNode{nil, nil, v}
	node := bst.root
	for {
		if v < node.value {
			if node.left == nil {
				node.left = newNode
				return
			}
			node = node.left
		} else {
			if node.right == nil {
				node.right = newNode
				return
			}
			node = node.right
		}
	}
}

func (bst *BinarySearchTree) Exists(v int) bool {
	node := bst.root
	for {
		if node == nil {
			return false
		}
		if v == node.value {
			return true
		}
		if v < node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
}

func (bst *BinarySearchTree) Sorted() []int {
	var res []int
	return bst.inorderTraversal(bst.root, res)
}

func (bst *BinarySearchTree) inorderTraversal(node *BSTNode, res []int) []int {
	if node == nil {
		return res
	}

	res = bst.inorderTraversal(node.left, res)
	res = append(res, node.value)
	res = bst.inorderTraversal(node.right, res)
	return res
}

func (bst *BinarySearchTree) SortedRev() []int {
	var res []int
	return bst.inorderTraversalRev(bst.root, res)
}

func (bst *BinarySearchTree) inorderTraversalRev(node *BSTNode, res []int) []int {
	if node == nil {
		return res
	}

	res = bst.inorderTraversalRev(node.right, res)
	res = append(res, node.value)
	res = bst.inorderTraversalRev(node.left, res)
	return res
}
