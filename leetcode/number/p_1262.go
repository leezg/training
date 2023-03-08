package number

func maxSumDivThree(nums []int) int {
	remain := make([]int, 3)
	for _, n := range nums {
		a := remain[0] + n
		b := remain[1] + n
		c := remain[2] + n
		remain[a%3] = max(remain[a%3], a)
		remain[b%3] = max(remain[b%3], b)
		remain[c%3] = max(remain[c%3], c)
	}

	return remain[0]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
