package array

func firstMissingPositive(nums []int) int {
	for i := range nums {
		if i >= len(nums) || i <= 0 {
			continue
		}
		for {
			if nums[i] >= len(nums) || nums[i] <= 0 {
				break
			}
			if nums[i] != nums[nums[i]-1] {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			} else {
				break
			}
		}
	}
	for i, num := range nums {
		if num-1 != i {
			return i + 1
		}
	}

	return len(nums) + 1
}
