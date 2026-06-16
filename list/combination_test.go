package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	input := [][]int{{1, 2}, {1, 2, 3}}
	expected := [][]int{{1, 1}, {1, 2}, {1, 3}, {2, 1}, {2, 2}, {2, 3}}

	got := Combinations(input)

	assert.Equal(t, expected, got)
}
