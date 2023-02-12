package string

import "strings"

func isPalindrome(str string) bool {
	s := strings.ToLower(str)
	ans := true
	i := 0
	j := len(s) - 1
	for i < j {
		if !isAlpha(s[j]) {
			j--
			continue
		}
		if !isAlpha(s[i]) {
			i++
			continue
		}
		if s[i] != s[j] {
			ans = false
			break
		}
		i++
		j--
	}
	return ans
}

func isAlpha(s uint8) bool {
	if s >= uint8('a') && s <= uint8('z') || s >= uint8('0') && s <= uint8('9') {
		return true
	}
	return false
}
