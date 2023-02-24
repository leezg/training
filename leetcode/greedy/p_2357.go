package greedy

func MinimumOperations(nums []int) int {
	return minimumOperations(nums)
}

func minimumOperations(nums []int) int {
	count := map[int]int{}
	for _, n := range nums {
		if n == 0 {
			continue
		}
		if _, ok := count[n]; !ok {
			count[n] = 1
		}
	}
	return len(count)
}
