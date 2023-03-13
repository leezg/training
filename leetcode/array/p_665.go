package array

func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			cnt++
			if cnt > 1 {
				return false
			}
			//出现1，2，3，9，4，5的情况不用该9的值，只是+1
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true
}
