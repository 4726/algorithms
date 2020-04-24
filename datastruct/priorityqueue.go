package datastruct

type PriorityQueue struct {
	heap *BinaryHeapMax
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		&BinaryHeapMax{},
	}
}

func (pq *PriorityQueue) Insert(v int) {
	pq.heap.Insert(v)
}

func (pq *PriorityQueue) Pop() int {
	if len(pq.heap.data) < 1 {
		return 0
	}
	highest := pq.heap.Max()
	heapLen := len(pq.heap.data)
	pq.heap.data[0], pq.heap.data[heapLen-1] = pq.heap.data[heapLen-1], pq.heap.data[0]
	pq.heap.data = pq.heap.data[:heapLen-1]
	pq.heap.heapifyDown(0)
	return highest
}

type BinaryHeapMax struct {
	data []int
}

func (bh *BinaryHeapMax) Insert(v int) {
	bh.data = append(bh.data, v)
	bh.heapifyUp(len(bh.data) - 1)
}

func (bh *BinaryHeapMax) heapifyUp(index int) {
	if index <= 0 {
		return
	}
	parentIndex := (index - 1) / 2
	parent := bh.data[parentIndex]
	if bh.data[index] > parent {
		bh.data[index], bh.data[parentIndex] = bh.data[parentIndex], bh.data[index]
		bh.heapifyUp(parentIndex)
	}
}

func (bh *BinaryHeapMax) heapifyDown(index int) {
	if len(bh.data) == 0 {
		return
	}
	leftIndex := (2 * index) + 1
	rightIndex := (2 * index) + 2
	greatest := bh.data[index]
	greatestIndex := index
	if len(bh.data) > leftIndex {
		left := bh.data[leftIndex]
		if left > greatest {
			greatest = left
			greatestIndex = leftIndex
		}
	}
	if len(bh.data) > rightIndex {
		right := bh.data[rightIndex]
		if right > greatest {
			greatest = right
			greatestIndex = rightIndex
		}
	}
	if index != greatestIndex {
		bh.data[index], bh.data[greatestIndex] = bh.data[greatestIndex], bh.data[index]
		bh.heapifyDown(greatestIndex)
	}
}

func (bh *BinaryHeapMax) Max() int {
	return bh.data[0]
}

func (bh *BinaryHeapMax) Heapsort(res []int) []int {
	if len(bh.data) < 1 {
		return res
	}

	res = append(res, bh.data[0])
	bh.data[0], bh.data[len(bh.data)-1] = bh.data[len(bh.data)-1], bh.data[0]
	bh.data = bh.data[:len(bh.data)-1]
	bh.heapifyDown(0)
	return bh.Heapsort(res)
}
