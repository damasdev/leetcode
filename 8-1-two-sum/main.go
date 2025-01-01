package main

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// TwoSum is a brute force solution for the TwoSum problem.
func TwoSum(nums []int, target int) []int {
	for i, num1 := range nums {
		for j, num2 := range nums {
			if i == j {
				continue
			}

			if num1+num2 == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// OptimizedTwoSum is an optimized solution for the TwoSum problem.
func OptimizedTwoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numsMap[target-num]; ok {
			return []int{j, i}
		}
		numsMap[num] = i
	}
	return nil
}
