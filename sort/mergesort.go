package sort

// in:  [7 3 8 4 1 5 5]
// in:  [7 3 8]
// in:  [7]
// in:  [3 8]
// in:  [3]
// in:  [8]
// out:  [3 8]
// out:  [3 7 8]
// in:  [4 1 5 5]
// in:  [4 1]
// in:  [4]
// in:  [1]
// out:  [1 4]
// in:  [5 5]
// in:  [5]
// in:  [5]
// out:  [5 5]
// out:  [1 4 5 5]
// out:  [1 3 4 5 5 7 8]

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
