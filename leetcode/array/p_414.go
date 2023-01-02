package array

import (
	"fmt"
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
		for _, a := range ans {
			if n == a {
				break
			} else if n > a {
				fmt.Println(ans)
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
