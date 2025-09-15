package main

func climbStairs(n int) int {
	a, b := 1, 0

	for range n {
		a, b = a+b, a
	}

	return a
}
