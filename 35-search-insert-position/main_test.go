package main_test

import (
	"testing"

	search "github.com/damasdev/leetcode/35-search-insert-position"
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
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := search.SearchInsert(tc.nums, tc.target)
			if got != tc.expected {
				t.Errorf("Search(%v, %d) = %d, want %d", tc.nums, tc.target, got, tc.expected)
			}
		})
	}
}
