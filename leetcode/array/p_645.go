package array

func FindErrorNums(nums []int) []int {
	return findErrorNums(nums)
}

func findErrorNums(nums []int) []int {
	n := len(nums)
	count := make([]int, n+1)
	sum := (n + 1) * n / 2
	eSum := 0
	eDouble := -1
	for _, x := range nums {
		eSum += x
		if count[x] != 0 {
			eDouble = x
		}
		count[x]++
	}
	eNull := sum - (eSum - eDouble)
	return []int{eDouble, eNull}
}
