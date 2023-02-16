package dp

var maxBonus []int

func rob(nums []int) int {
	maxBonus = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		maxBonus[i] = -1
	}
	return maxRob(len(nums)-1, nums)
}

func maxRob(n int, nums []int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return nums[0]
	}
	if maxBonus[n] != -1 {
		return maxBonus[n]
	}
	tmp := nums[n] + maxRob(n-2, nums)
	if tmp > maxBonus[n] {
		maxBonus[n] = tmp
	}
	tmp = maxRob(n-1, nums)
	if tmp > maxBonus[n] {
		maxBonus[n] = tmp
	}
	return maxBonus[n]
}
