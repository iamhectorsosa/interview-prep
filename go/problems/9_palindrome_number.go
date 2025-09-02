package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	c := x
	r := 0
	for c > 0 {
		r = r*10 + (c % 10)
		c /= 10
	}
	return x == r
}

func isPalindromeStr(x int) bool {
	s := strconv.Itoa(x)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return s == string(runes)
}
