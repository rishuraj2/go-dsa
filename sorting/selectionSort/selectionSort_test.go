package selectionsort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedOutput []int
	} {
		{"#1", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"#2", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"#3", []int{1}, []int{1}},
		{"#4", []int{}, []int{}},
		{"#5", []int{58, 6, 34, 2, 23, 23, 34, 43, 65, 75}, []int{2, 6, 23, 23, 34, 34, 43, 58, 65, 75}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inputArray := append([]int{}, test.input...)
			SelectionSort(inputArray);
			if !reflect.DeepEqual(inputArray, test.expectedOutput) {
				t.Errorf("expected: %v, but got: %v", test.expectedOutput, inputArray);
			}
		})
	}
}
