package array

import (
	"math"
)

func ThirdMax(nums []int) int {
	return thirdMax(nums)
}

func thirdMax(nums []int) int {
	var ans []int
	for i := 0; i < 3; i++ {
		ans = append(ans, math.MinInt)
	}
	for _, n := range nums {
		for i, a := range ans {
			if n == a {
				break
			} else if n > a {
				tmp := append(ans[:i+1], ans[i:]...)
				tmp[i] = n
				ans = tmp[:3]
				break
			}

		}
	}
	if ans[2] == math.MinInt {
		return ans[0]
	} else {
		return ans[2]
	}
}
