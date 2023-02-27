package array

func movesToMakeZigzag(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	count1 := 0
	count2 := 0
	nums2 := make([]int, len(nums))
	for i, n := range nums {
		nums2[i] = n
	}
	for i := 0; i < len(nums)-1; i++ {
		if i%2 == 0 {
			if nums[i] <= nums[i+1] {
				count1 += nums[i+1] - nums[i] + 1
				nums[i+1] = nums[i] - 1
			}
			if nums2[i] >= nums2[i+1] {
				count2 += nums2[i] - nums2[i+1] + 1
				nums2[i] = nums2[i+1] - 1
			}
		} else {
			if nums[i] >= nums[i+1] {
				count1 += nums[i] - nums[i+1] + 1
				nums[i] = nums[i+1] - 1
			}
			if nums2[i] <= nums[i+1] {
				count2 += nums2[i+1] - nums2[i] + 1
				nums2[i+1] = nums2[i] - 1
			}
		}
	}
	if count2 < count1 {
		return count2
	}
	return count1
}
