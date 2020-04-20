package sort

// in:  [7 3 8 4 1 5 5]
// in:  [3 4 1 5 5]
// in:  [1]
// in:  [4 5 5]
// in:  []
// in:  [5 5]
// in:  []
// in:  [5]
// out:  [5 5]
// out:  [4 5 5]
// out:  [1 3 4 5 5]
// in:  [8]
// out:  [1 3 4 5 5 7 8]

func Quicksort(in []int) {
	if len(in) <= 1 {
		return
	}

	var left, right []int
	pivot := in[0]
	for i := 1; i < len(in); i++ {
		if in[i] < pivot {
			left = append(left, in[i])
		} else {
			right = append(right, in[i])
		}
	}
	Quicksort(left)
	Quicksort(right)
	copy(in, left)
	in[len(left)] = pivot
	copy(in[len(left)+1:], right)
}
