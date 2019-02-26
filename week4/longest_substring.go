package main

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	filter := make(map[rune]string, 0)
	l := 0

	for _, char := range s {
		if _, exists := filter[char]; exists {
			if len(filter) > l {
				l = len(filter)
			}
			filter = make(map[rune]string, 0)
		}
		filter[char] = ""
	}

	if len(filter) > l {
		return len(filter)
	}

	return l
}
