package largestelementinanarray

import (
	"testing"
)

func TestLargestElementInAnAray(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		expectedValue int
	}{
		{"Empty", []int{}, 0},
		{"Normal", []int{1, 3, 5, 2}, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := LargestElementInAnArray(test.input)

			if result != test.expectedValue {
				t.Errorf("expected: %d but got %d", test.expectedValue, result)
			}
		})
	}
}
