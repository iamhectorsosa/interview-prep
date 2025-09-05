package main

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []rune

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) < 1 || pairs[char] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}
