package removeduplicatesfromsortedarray

import (
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedOutput int
	} {
		{"#1", []int{1, 1, 2}, 2},
		{"#2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RemoveDuplicatesFromSortedArray(test.input)

			if result != test.expectedOutput {
				t.Errorf("expected %d but got %d", test.expectedOutput, result)
			}
		})
	}
}
