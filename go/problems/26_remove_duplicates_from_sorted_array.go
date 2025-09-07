package main

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	hashMap := make(map[int]struct{})
	writeIndex := 0
	for _, num := range nums {
		if _, ok := hashMap[num]; !ok {
			hashMap[num] = struct{}{}
			nums[writeIndex] = num
			writeIndex++
		}
	}
	return len(hashMap)
}
