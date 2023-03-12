package array

func FindDuplicates(nums []int) []int {
	return findDuplicates(nums)
}

//func findDuplicates(nums []int) []int {
//	index := 0
//	for _, n := range nums {
//		index = index ^ (1 << (n - 1))
//	}
//	for i := 0; i < len(nums); i++ {
//		if index&(1<<(nums[i]-1)) == 0 {
//			index = index ^ (1 << (nums[i] - 1))
//		} else {
//			nums = append(nums[:i], nums[i+1:]...)
//			i--
//		}
//	}
//	return nums
//}

//注意n的大小

func findDuplicates(nums []int) (ans []int) {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num-1 != i {
			ans = append(ans, num)
		}
	}
	return
}
