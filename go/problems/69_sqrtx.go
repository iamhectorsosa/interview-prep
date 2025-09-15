package main

func mySqrt(x int) int {
	// sqrt := 0
	// for num, oddNum := x, 1; num >= oddNum; oddNum += 2 {
	// 	sqrt++
	// 	num -= oddNum
	// }
	//
	// return sqrt

	l := 0
	r := x

	for l < r {
		m := (l + r + 1) / 2
		if m*m <= x {
			l = m
			continue
		}
		r = m - 1
	}

	return l
}
