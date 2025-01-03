package main_test

import (
	"testing"

	search "github.com/damasdev/leetcode/704-binary-search"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
		{
			name:     "Single element found",
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element not found",
			nums:     []int{5},
			target:   1,
			expected: -1,
		},
		{
			name:     "Target at beginning",
			nums:     []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "Target at end",
			nums:     []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			target:   1,
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := search.Search(tc.nums, tc.target)
			if got != tc.expected {
				t.Errorf("Search(%v, %d) = %d, want %d", tc.nums, tc.target, got, tc.expected)
			}
		})
	}
}
