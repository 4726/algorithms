package sort

func Bubblesort(in []int) {
	var sorted bool
	for !sorted {
		sorted = true
		for i := 0; i < len(in)-1; i++ {
			if in[i] > in[i+1] {
				in[i+1], in[i] = in[i], in[i+1]
				sorted = false
			}
		}
	}
}
