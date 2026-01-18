package checkifarrayissortedandrotated

import (
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedOutput bool
	} {
		{"#1", []int{}, false},
		{"#2", []int{3, 4, 5, 1, 2}, true},
		{"#3", []int{3, 4, 5, 1, 2}, true},
		{"#4", []int{2, 1, 3, 4}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := Check(test.input)

			if output != test.expectedOutput {
				t.Errorf("expected: %t but got %t", test.expectedOutput, output)
			}
		})
	}
}
