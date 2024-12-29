package search_test

import (
	"testing"

	search "github.com/damasdev/leetcode/704-binary-search"
)

// TestSearch tests the Search function from the search package
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
