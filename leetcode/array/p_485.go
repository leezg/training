package array

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	tmp := 0
	for _, n := range nums {
		if n == 1 {
			tmp++
		} else if tmp != 0 {
			if tmp > count {
				count = tmp
			}
			tmp = 0
		}
	}
	if tmp > count {
		count = tmp
	}
	return count
}
