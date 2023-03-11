package array

func FindLongestSubarray(array []string) []string {
	return findLongestSubarray(array)
}

func findLongestSubarray(array []string) []string {
	m := map[int]int{0: -1}
	sum := 0
	l, r := 0, 0
	for i, n := range array {

		if n[0] >= '0' && n[0] <= '9' {
			sum++
		} else {
			sum--
		}
		if index, ok := m[sum]; ok {
			if i-index > r-l {
				r = i + 1
				l = index + 1
			}
		} else {
			m[sum] = i
		}
	}
	return array[l:r]
}
