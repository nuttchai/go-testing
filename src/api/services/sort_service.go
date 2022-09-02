package services

import "github.com/nuttchai/go-testing/src/api/utils/sort"

const (
	privateConst = "private"
	PublicConst  = "public"
)

func Sort(elements []int) {
	// smaller number of elements => BubbleSort is faster
	if len(elements) <= 20000 {
		sort.BubbleSort(elements)
		return
	}

	// however if number of elements get increased => Sort from go library is much faster
	sort.Sort(elements)
}
