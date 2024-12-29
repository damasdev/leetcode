package main

import "fmt"

// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
// If target exists, then return its index. Otherwise, return -1.
// You must write an algorithm with O(log n) runtime complexity.
func search(nums []int, target int) int {
	half := len(nums) / 2
	if nums[half] == target {
		return half
	}

	if target < nums[half] {
		for i := half; i >= 0; i-- {
			if nums[i] == target {
				return i
			}
		}
	} else {
		for i := half; i < len(nums); i++ {
			if nums[i] == target {
				return i
			}
		}
	}

	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	val := search(nums, target)
	fmt.Println(val)
}
