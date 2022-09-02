package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/nuttchai/go-testing/src/api/services"
)

// NOTE: We can separate the test file to different package (it would be black box testing because the package is outside). However, it would be more recommended to have white box testing to be able to test private function and constant like privateConst
// func TestConstant(t *testing.T) {
//	// NOTE: cannot be tested because it is a private constant that is not exported
// 	if privateConst != "private" {
// 		t.Error("privateConst should be private")
// 	}
// 	if PublicConst != "public" {
// 		t.Error("privateConst should be public")
// 	}
// }

func TestSort(t *testing.T) {
	elements := []int{7, 9, 5, 3, 1, 2, 8, 4, 6}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	services.Sort(elements)

	if !reflect.DeepEqual(expected, elements) {
		fmt.Printf("expected: %v\n", expected)
		fmt.Printf("result: %v\n", elements)
		t.Error("Expected from BubbleSort result is incorrect")
	}
}
