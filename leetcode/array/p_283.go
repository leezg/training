package array

func MoveZeroes(nums []int) {
	moveZeroes(nums)
}

//func moveZeroes(nums []int) {
//	n := len(nums)
//	for i := 0; i < n; i++ {
//		for nums[i] == 0 && i < n {
//			for j := i; j < n-1; j++ {
//				nums[j], nums[j+1] = nums[j+1], nums[j]
//			}
//			n--
//		}
//	}
//}

func moveZeroes(nums []int) {
	i, j := 0, 0
	for i < len(nums) && nums[i] != 0 {
		i++
	}
	j = i
	for i++; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}
