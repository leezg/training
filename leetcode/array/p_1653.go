package array

import "strings"

func minimumDeletions(s string) int {
	var res int
	aTotal := strings.Count(s, "a")
	bTotal := len(s) - aTotal
	aCount := 0
	bCount := 0
	if aTotal < bTotal {
		res = aTotal
	} else {
		res = bTotal
	}
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			aCount++
		} else {
			bCount++
		}
		if bCount+aTotal-aCount < res {
			res = bCount + aTotal - aCount
		}
	}
	return res
}
