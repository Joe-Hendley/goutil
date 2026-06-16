package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	for _, tc := range []struct {
		input                [][]int
		expectedCombinations [][]int
	}{
		{
			input: [][]int{{1, 2}, {1, 2, 3}},
			expectedCombinations: [][]int{
				{1, 1}, {1, 2}, {1, 3},
				{2, 1}, {2, 2}, {2, 3},
			},
		},
		{
			input: [][]int{{1, 2}, {3, 4}, {5, 6}},
			expectedCombinations: [][]int{
				{1, 3, 5}, {1, 3, 6}, {1, 4, 5}, {1, 4, 6},
				{2, 3, 5}, {2, 3, 6}, {2, 4, 5}, {2, 4, 6},
			},
		}, {
			input: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			expectedCombinations: [][]int{
				{1, 3, 5, 7}, {1, 3, 6, 7}, {1, 4, 5, 7}, {1, 4, 6, 7},
				{2, 3, 5, 7}, {2, 3, 6, 7}, {2, 4, 5, 7}, {2, 4, 6, 7},
				{1, 3, 5, 8}, {1, 3, 6, 8}, {1, 4, 5, 8}, {1, 4, 6, 8},
				{2, 3, 5, 8}, {2, 3, 6, 8}, {2, 4, 5, 8}, {2, 4, 6, 8},
			},
		},
		{
			input: [][]int{{1}, {2}, {3}, {4}},
			expectedCombinations: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			input: [][]int{{1}, {2, 3}},
			expectedCombinations: [][]int{
				{1, 2}, {1, 3},
			},
		},
		{
			input: [][]int{{1}, {2}, {3, 4}},
			expectedCombinations: [][]int{
				{1, 2, 3}, {1, 2, 4},
			},
		},
		{
			input: [][]int{{1}, {2}, {3}, {4, 5}},
			expectedCombinations: [][]int{
				{1, 2, 3, 4}, {1, 2, 3, 5},
			},
		},
		{
			input: [][]int{{1, 2}, {3}, {4}, {5}},
			expectedCombinations: [][]int{
				{1, 3, 4, 5}, {2, 3, 4, 5},
			},
		},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			got := Combinations(tc.input)

			assert.Len(t, got, len(tc.expectedCombinations))

			for _, combination := range tc.expectedCombinations {
				assert.Contains(t, got, combination)
			}
		})

	}
}
