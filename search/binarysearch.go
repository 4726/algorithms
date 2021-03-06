package search

func Binarysearch(in []int, value int) int {
	mid := in[len(in)/2]
	if value == mid {
		return len(in) / 2
	}

	if len(in) == 1 {
		return -1
	}

	if value < mid {
		return Binarysearch(in[:len(in)/2], value)
	} else {
		res := Binarysearch(in[len(in)/2:], value)
		if res == -1 {
			return -1
		}
		return res + len(in)/2
	}
}
