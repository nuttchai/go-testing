package services

import (
	"fmt"
	"reflect"
	"testing"
)

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
		t.Error("Expected from BubbleSort result is incorrect")
	}
}
