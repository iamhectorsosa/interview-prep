package main

func climbStairs(n int) int {
	a, b := 0, 1

	for range n {
		a, b = b, a+b
	}

	return b
}
