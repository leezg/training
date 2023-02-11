package array

func FindDisappearedNumbers(nums []int) []int {
	return findErrorNums(nums)
}

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	label := make([]int, n+1)
	label[0] = 1
	for _, x := range nums {
		label[x] = 1
	}
	var ans []int
	for i, n := range label {
		if n == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}
