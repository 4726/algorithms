package sort

func Insertionsort(in []int) {
	for i := 0; i < len(in); i++ {
		num := in[i]
		var found bool
		for j := i - 1; j >= 0; j-- {
			if num < in[j] {
				in[j+1] = in[j]
			} else {
				in[j+1] = num
				found = true
				break
			}
		}
		if !found {
			in[0] = num
		}
	}
}
