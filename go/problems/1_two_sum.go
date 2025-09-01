package main

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if key, exists := hashMap[diff]; exists {
			return []int{key, i}
		}
		hashMap[num] = i
	}
	return []int{}
}
