package sort

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
