package sort

func Mergesort(in []int) {
	if len(in) <= 1 {
		return
	}

	left := in[:len(in)/2]
	right := in[len(in)/2:]
	Mergesort(left)
	Mergesort(right)
	var sorted []int
	var leftIndex, rightIndex int
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			sorted = append(sorted, left[leftIndex])
			leftIndex++
		} else {
			sorted = append(sorted, right[rightIndex])
			rightIndex++
		}
	}
	if leftIndex < len(left) {
		sorted = append(sorted, left[leftIndex:]...)
	} else if rightIndex < len(right) {
		sorted = append(sorted, right[rightIndex:]...)
	}
	copy(in, sorted)
}
