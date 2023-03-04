package bit

func countTriplets(nums []int) int {
	var cnt [1 << 16]int
	for i := range nums {
		for j := range nums {
			cnt[nums[i]&nums[j]]++
		}
	}
	res := 0
	for i := range nums {
		x := nums[i] ^ 0xffff
		for sub := x; sub > 0; sub = (sub - 1) & x {
			res += cnt[sub]
		}
		res += cnt[0]
	}
	return res
}
