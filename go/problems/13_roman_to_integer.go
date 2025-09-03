package main

func romanToInt(s string) int {
	if len(s) < 1 || len(s) > 15 {
		return 0
	}

	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	runes := []rune(s)
	res := 0

	for i, r := range runes {
		val, ok := values[r]
		if !ok {
			return 0
		}

		if i+1 < len(runes) && values[runes[i+1]] > val {
			res -= val
			continue
		}

		res += val
	}

	return res
}
