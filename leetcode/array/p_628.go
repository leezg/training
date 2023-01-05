package array

import "math"

func maximumProduct(nums []int) int {
	var max3 []int
	var min2 []int
	for i := 0; i < 3; i++ {
		max3 = append(max3, math.MinInt)
	}
	for i := 0; i < 2; i++ {
		min2 = append(min2, math.MaxInt)
	}
	for _, n := range nums {
		for i, a := range max3 {
			if n > a {
				tmp := append(max3[:i+1], max3[i:]...)
				tmp[i] = n
				max3 = tmp[:3]
				break
			}
		}
		for i, a := range min2 {
			if n < a {
				tmp := append(min2[:i+1], min2[i:]...)
				tmp[i] = n
				min2 = tmp[:2]
				break
			}
		}
	}
	ans := max3[0] * max3[1] * max3[2]
	if min2[1] < 0 {
		tmp := max3[0] * min2[0] * min2[1]
		if tmp > ans {
			ans = tmp
		}
	}
	return ans
}
