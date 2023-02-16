package array

func numberOfPairs(nums []int) []int {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
	ForLoop1:
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				nums = append(nums[:j], nums[j+1:]...)
				nums = append(nums[:i], nums[i+1:]...)
				i--
				count++
				break ForLoop1
			}
		}
	}
	return []int{count, len(nums)}
}
