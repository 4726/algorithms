package sort

func Selectionsort(in []int) {
	for i := 0; i < len(in); i++ {
		least := in[i]
		leastIndex := i
		for j := i; j < len(in); j++ {
			if in[j] < least {
				least = in[j]
				leastIndex = j
			}
		}
		in[i], in[leastIndex] = in[leastIndex], in[i]
	}
}
