package elements

func GetElements(n int) []int {
	result := make([]int, n)
	j := 0

	for i := n; i >= 1; i-- {
		result[j] = i
		j++
	}
	return result
}
