package main

func removeElement(nums []int, val int) int {
	writeIndex := 0
	for _, num := range nums {
		if num != val {
			nums[writeIndex] = num
			writeIndex++
		}
	}

	nums = nums[:writeIndex]
	return len(nums)
}
