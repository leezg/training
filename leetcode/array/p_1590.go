package array

func minSubarray(nums []int, p int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	rem := sum % p
	if rem == 0 {
		return 0
	}

	sum = 0
	res := len(nums)
	m := map[int]int{0: -1}
	for i, n := range nums {
		sum += n
		tempRem := sum % p
		if index, ok := m[(tempRem-rem+p)%p]; ok {
			res = min(res, i-index)
		}
		m[tempRem%p] = i
	}
	if res >= len(nums) {
		return -1
	}
	return res
}
