package services

import (
	"fmt"
	"reflect"
	"testing"
)

func getElements(n int) []int {
	result := make([]int, n)
	j := 0

	for i := n; i > 1; i-- {
		result[j] = i
		j++
	}
	return result
}

// NOTE: If we only run TestConstant function, it will give 0% coverage
// However, it is still a good practice to test a constant
func TestConstant(t *testing.T) {
	if privateConst != "private" {
		t.Error("privateConst should be private")
	}
	if PublicConst != "public" {
		t.Error("privateConst should be public")
	}
}

func TestSort(t *testing.T) {
	elements := []int{7, 9, 5, 3, 1, 2, 8, 4, 6}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	Sort(elements)

	if !reflect.DeepEqual(expected, elements) {
		fmt.Printf("expected: %v\n", expected)
		fmt.Printf("result: %v\n", elements)
		t.Error("Expected from Sort result is incorrect")
	}
}

func TestSortMoreThan10k(t *testing.T) {
	number := 10000
	elements := getElements(number)

	Sort(elements)

	if elements[number-1] != number {
		fmt.Printf("expected: %v\n", number)
		fmt.Printf("result: %v\n", elements[number-1])
		t.Error("Expected from Sort result is incorrect")
	}
	if elements[0] != 0 {
		fmt.Printf("expected: %v\n", 0)
		fmt.Printf("result: %v\n", elements[0])
		t.Error("Expected from Sort result is incorrect")
	}
}

func TestSortMoreThan30k(t *testing.T) {
	number := 30000
	elements := getElements(number)

	Sort(elements)

	if elements[number-1] != number {
		fmt.Printf("expected: %v\n", number)
		fmt.Printf("result: %v\n", elements[number-1])
		t.Error("Expected from Sort result is incorrect")
	}
	if elements[0] != 0 {
		fmt.Printf("expected: %v\n", 0)
		fmt.Printf("result: %v\n", elements[0])
		t.Error("Expected from Sort result is incorrect")
	}
}

func BenchmarkSortMoreThan20k(b *testing.B) {
	number := 20000
	elements := getElements(number)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
