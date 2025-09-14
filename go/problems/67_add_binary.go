package main

import (
	"strconv"
)

func addBinary(a string, b string) string {
	carry := 0
	result := ""
	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		carry = sum / 2
		char := sum % 2
		result = strconv.Itoa(char) + result
	}

	if carry > 0 {
		return "1" + result
	}
	return result
}
