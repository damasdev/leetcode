package main_test

// Package main_test provides test cases for the Two Sum problem solutions.
// It tests both the basic and optimized implementations of finding two numbers
// that add up to a target value in a given array.

import (
	"testing"

	twosum "github.com/damasdev/leetcode/8-1-two-sum"
)

// testCases contains the test scenarios for verifying the Two Sum implementations.
// Each test case includes a name, input array, target sum, and expected output indices.
var testCases = []struct {
	name     string
	nums     []int
	target   int
	expected []int
}{
	{
		name:     "Example 1",
		nums:     []int{2, 7, 11, 15},
		target:   9,
		expected: []int{0, 1},
	},
	{
		name:     "Example 2",
		nums:     []int{3, 2, 4},
		target:   6,
		expected: []int{1, 2},
	},
	{
		name:     "Example 3",
		nums:     []int{3, 3},
		target:   6,
		expected: []int{0, 1},
	},
}

// TestTwoSum tests the basic implementation of the Two Sum solution.
// It verifies that the function correctly returns indices of two numbers
// that sum to the target value.
func TestTwoSum(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := twosum.TwoSum(tc.nums, tc.target)
			if len(got) != 2 {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.expected)
			}

			if got[0] != tc.expected[0] || got[1] != tc.expected[1] {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.expected)
			}
		})
	}
}

// TestOptimizedTwoSum tests the optimized implementation of the Two Sum solution.
// It verifies that the function correctly returns indices of two numbers
// that sum to the target value using a more efficient algorithm.
func TestOptimizedTwoSum(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := twosum.OptimizedTwoSum(tc.nums, tc.target)
			if len(got) != 2 {
				t.Errorf("OptimizedTwoSum(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.expected)
			}

			if got[0] != tc.expected[0] || got[1] != tc.expected[1] {
				t.Errorf("OptimizedTwoSum(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.expected)
			}
		})
	}
}
