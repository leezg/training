package string

func longestCommonPrefix(strs []string) string {
	a := strs[0]
	for i := 1; i < len(strs); i++ {
		a = findPrefix(a, strs[i])
		if len(a) == 0 {
			return ""
		}
	}
	return a
}

func findPrefix(a, b string) string {
	i := 0
	for ; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[:i]
		}
	}
	return a[:i]
}
