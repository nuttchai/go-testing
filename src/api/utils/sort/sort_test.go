package sort

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/nuttchai/go-testing/src/api/utils/elements"
	"github.com/stretchr/testify/assert"
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

// NOTE: Test function with asset (another way to test unit test)
// If the a part of the function starts to fail, the test will immediately stop at that point, instead of continue running the remaining tests of the function like above)
func TestBubbleSortWithAsset(t *testing.T) {
	elements := elements.GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 10, elements[0])
	assert.EqualValues(t, 1, elements[len(elements)-1])

	BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 1, elements[0], "first element should be 1")
	assert.EqualValues(t, 10, elements[len(elements)-1], "last element should be 10")
}
