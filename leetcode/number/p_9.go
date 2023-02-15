package number

import "fmt"

func isPalindrome(x int) bool {
	s := fmt.Sprint(x)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
