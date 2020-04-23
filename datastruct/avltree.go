package datastruct

type AVLTree struct {
	root *AVLNode
}

type AVLNode struct {
	left, right, parent *AVLNode
	value               int
	height              int
}

func (a *AVLTree) Insert(v int) {
	newNode := &AVLNode{nil, nil, nil, v, 0}
	if a.root == nil {
		a.root = newNode
		return
	}

	node := a.root
	for {
		if v < node.value {
			if node.left == nil {
				node.left = newNode
				newNode.parent = node
				a.updateParentHeights(node)
				a.balance(node, newNode)
				return
			}
			node = node.left
		} else {
			if node.right == nil {
				node.right = newNode
				newNode.parent = node
				a.updateParentHeights(node)
				a.balance(node, newNode)
				return
			}
			node = node.right
		}
	}

}

func (a *AVLTree) updateParentHeights(n *AVLNode) {
	if n == nil {
		return
	}
	a.updateNodeHeight(n)
	a.updateParentHeights(n.parent)
}

func (a *AVLTree) updateNodeHeight(n *AVLNode) {
	if n == nil {
		return
	}
	var nLeftHeight int
	var nRightHeight int
	if n.left == nil {
		nLeftHeight = -1
	} else {
		nLeftHeight = n.left.height
	}
	if n.right == nil {
		nRightHeight = -1
	} else {
		nRightHeight = n.right.height
	}
	if nLeftHeight > nRightHeight {
		n.height = nLeftHeight + 1
	} else {
		n.height = nRightHeight + 1
	}
}

func (a *AVLTree) balance(n, added *AVLNode) {
	if n == nil {
		return
	}
	var nLeftHeight int
	var nRightHeight int
	if n.left == nil {
		nLeftHeight = -1
	} else {
		nLeftHeight = n.left.height
	}
	if n.right == nil {
		nRightHeight = -1
	} else {
		nRightHeight = n.right.height
	}
	balance := nRightHeight - nLeftHeight
	if balance < -1 {
		z := n
		y := z.left
		if added.value < y.value {
			a.rightRotate(z)
		} else {
			a.leftRotate(y)
			a.rightRotate(z)
		}
		return
	}
	if balance > 1 {
		z := n
		y := z.right
		if added.value < y.value {
			a.rightRotate(y)
			a.leftRotate(z)
		} else {
			a.leftRotate(z)
		}
		return
	}
	a.balance(n.parent, added)
}

func (a *AVLTree) leftRotate(n *AVLNode) {
	z := n
	y := z.right

	z.right = y.left
	if a.root == z {
		a.root = y
	} else {
		if z.value < z.parent.value {
			z.parent.left = y
		} else {
			z.parent.right = y
		}
	}
	z.parent = y
	y.left = z
	a.updateNodeHeight(z)
	a.updateNodeHeight(y)
}

func (a *AVLTree) rightRotate(n *AVLNode) {
	z := n
	y := z.left

	z.left = y.right
	if a.root == z {
		a.root = y
	} else {
		if z.value < z.parent.value {
			z.parent.left = y
		} else {
			z.parent.right = y
		}
	}
	z.parent = y
	y.right = z
	a.updateNodeHeight(z)
	a.updateNodeHeight(y)
}
