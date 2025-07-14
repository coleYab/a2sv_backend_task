package bb

import (
	"strings"
	"unicode"
)

func CheckPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	s = strings.ToLower(s)

	isLetter := func(a rune) bool {
		return unicode.IsLetter(a)
	}

	for left < right {
		if !isLetter(rune(s[left])) {
			left++
			continue
		}
		if !isLetter(rune(s[right])) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}
	}

	return true
}
