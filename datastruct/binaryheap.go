package datastruct

type BinaryHeap struct {
	data []int
}

func (bh *BinaryHeap) Insert(v int) {
	bh.data = append(bh.data, v)
	bh.heapifyUp(len(bh.data) - 1)
}

func (bh *BinaryHeap) heapifyUp(index int) {
	if index <= 0 {
		return
	}
	parentIndex := (index - 1) / 2
	parent := bh.data[parentIndex]
	if bh.data[index] < parent {
		bh.data[index], bh.data[parentIndex] = bh.data[parentIndex], bh.data[index]
		bh.heapifyUp(parentIndex)
	}
}

func (bh *BinaryHeap) heapifyDown(index int) {
	if len(bh.data) == 0 {
		return
	}
	leftIndex := (2 * index) + 1
	rightIndex := (2 * index) + 2
	least := bh.data[index]
	leastIndex := index
	if len(bh.data) > leftIndex {
		left := bh.data[leftIndex]
		if left < least {
			least = left
			leastIndex = leftIndex
		}
	}
	if len(bh.data) > rightIndex {
		right := bh.data[rightIndex]
		if right < least {
			least = right
			leastIndex = rightIndex
		}
	}
	if index != leastIndex {
		bh.data[index], bh.data[leastIndex] = bh.data[leastIndex], bh.data[index]
		bh.heapifyDown(leastIndex)
	}
}

func (bh *BinaryHeap) Min() int {
	return bh.data[0]
}

func (bh *BinaryHeap) Heapsort(res []int) []int {
	if len(bh.data) < 1 {
		return res
	}

	res = append(res, bh.data[0])
	bh.data[0], bh.data[len(bh.data)-1] = bh.data[len(bh.data)-1], bh.data[0]
	bh.data = bh.data[:len(bh.data)-1]
	bh.heapifyDown(0)
	return bh.Heapsort(res)
}
