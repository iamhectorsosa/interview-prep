package main

func lengthOfLastWord(s string) int {
	// if len(s) == 0 {
	// 	return 0
	// }
	//
	// var length int
	// for i := len(s) - 1; i >= 0; i-- {
	// 	if string(s[i]) != " " {
	// 		length++
	// 		continue
	// 	}
	// 	if length > 0 && string(s[i]) == " " {
	// 		return length
	// 	}
	// }
	//
	// return length

	end := len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}

	if end < 0 {
		return 0
	}

	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}

	return end - start
}
