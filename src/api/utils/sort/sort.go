package sort

func BubbleSort(elements []int) {
	keepSorting := true
	for keepSorting {
		keepSorting = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				keepSorting = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}
