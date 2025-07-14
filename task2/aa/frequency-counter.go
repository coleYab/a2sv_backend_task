package aa

func CountFrequncy(s string) map[string]int {
	res := map[string]int{}

	left := 0
	for right := 0; right <= len(s); right++ {
		isLetterOrAppostoph := s[right] != '\'' && (('a' <= s[right] && 'z' >= s[right]) || ('A' <= s[right] && 'Z' >= s[right]))
		if !isLetterOrAppostoph {
			subStr := s[left:right]
			if subStr != "" {
				res[subStr] += 1
			}
			left = right + 1
		}

		if right == len(s)-1 {
			subStr := s[left : right+1]
			if subStr != "" {
				res[subStr] += 1
			}
		}
	}

	return res
}
