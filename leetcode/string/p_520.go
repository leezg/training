package string

import "regexp"

func DetectCapitalUse(word string) bool {
	return detectCapitalUse(word)
}

func detectCapitalUse(word string) bool {
	rules := "^[A-Z]+$|^[a-z]+$|^[A-Z][a-z]+$"
	if word != "" {
		if ok, _ := regexp.MatchString(rules, word); ok {
			return true
		}
	}
	return false
}
