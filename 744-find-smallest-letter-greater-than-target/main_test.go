package main_test

import (
	"testing"

	search "github.com/damasdev/leetcode/744-find-smallest-letter-greater-than-target"
)

func TestNextGreatestLetter(t *testing.T) {
	testCases := []struct {
		name     string
		letters  []byte
		target   byte
		expected byte
	}{
		{
			name:     "Example 1",
			letters:  []byte{'c', 'f', 'j'},
			target:   'a',
			expected: 'c',
		},
		{
			name:     "Example 2",
			letters:  []byte{'c', 'f', 'j'},
			target:   'c',
			expected: 'f',
		},
		{
			name:     "Example 3",
			letters:  []byte{'x', 'x', 'y', 'y'},
			target:   'z',
			expected: 'x',
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := search.NextGreatestLetter(tc.letters, tc.target)
			if got != tc.expected {
				t.Errorf("NextGreatestLetter(%v, %c) = %c, want %c", tc.letters, tc.target, got, tc.expected)
			}
		})
	}
}
