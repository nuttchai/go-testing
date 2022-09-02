package sort

import (
	"fmt"
	"reflect"
	"testing"
)

// NOTE: any test MUST starts with a word "Test"
func TestBubbleSort(t *testing.T) {
	// Step 1: Init (Optional)
	expected := []int{0, 1, 5, 6, 7, 8, 9, 9}
	elements := []int{9, 5, 6, 1, 7, 8, 0, 9}

	// Step 2: Execution
	BubbleSort(elements)

	// Step 3: Validation
	if !reflect.DeepEqual(expected, elements) {
		fmt.Printf("expected: %v\n", expected)
		fmt.Printf("result: %v\n", elements)
		t.Error("Expected from BubbleSort result is incorrect")
	}
}

// NOTE: if we only test this function, this would not give 100% coverage because it didn't run its swapping logic part
func TestBubbleSortWithAlreadySortedInput(t *testing.T) {
	// Step 1: Init (Optional)
	expected := []int{0, 1, 5, 6, 7, 8, 9, 9}
	elements := []int{0, 1, 5, 6, 7, 8, 9, 9} // same with expected

	// Step 2: Execution
	BubbleSort(elements)

	// Step 3: Validation
	if !reflect.DeepEqual(expected, elements) {
		fmt.Printf("expected: %v\n", expected)
		fmt.Printf("result: %v\n", elements)
		t.Error("Expected from BubbleSort result is incorrect")
	}
}

func TestSort(t *testing.T) {
	// Step 1: Init (Optional)
	expected := []int{0, 1, 5, 6, 7, 8, 9, 9}
	elements := []int{9, 5, 6, 1, 7, 8, 0, 9}

	// Step 2: Execution
	Sort(elements)

	// Step 3: Validation
	if !reflect.DeepEqual(expected, elements) {
		fmt.Printf("expected: %v\n", expected)
		fmt.Printf("result: %v\n", elements)
		t.Error("Expected from BubbleSort result is incorrect")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := []int{9, 5, 6, 1, 7, 8, 0, 9}

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := []int{9, 5, 6, 1, 7, 8, 0, 9}

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
