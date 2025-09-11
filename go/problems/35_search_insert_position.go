package main

func searchInsert(nums []int, target int) int {
	leftBoundary := 0
	rightBoundary := len(nums) - 1

	for leftBoundary <= rightBoundary {
		middle := (leftBoundary + rightBoundary) / 2
		if target > nums[middle] {
			leftBoundary = middle + 1
		} else {
			rightBoundary = middle - 1
		}
	}
	return leftBoundary
}
