package sort

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/nuttchai/go-testing/src/api/utils/elements"
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

func TestBubbleSortTimeout(t *testing.T) {
	elements := elements.GetElements(10)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	timeoutChan := make(chan bool, 1)

	// NOTE: this goroutine runs BubbleSort
	go func() {
		BubbleSort(elements)
		timeoutChan <- false
	}()

	// NOTE: this goroutine counts timeout
	go func() {
		time.Sleep(500 * time.Millisecond)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		t.Error("BubbleSort took more than 500 millisecond")
	}

	if !reflect.DeepEqual(expected, elements) {
		fmt.Printf("expected: %v\n", expected)
		fmt.Printf("result: %v\n", elements)
		t.Error("Expected from BubbleSort result is incorrect")
	}
}
