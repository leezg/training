package _2ptr

//func removeElement(nums []int, val int) int {
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == val {
//			nums = append(nums[:i], nums[i+1:]...)
//			i--
//		}
//	}
//	return len(nums)
//}
//better:
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
