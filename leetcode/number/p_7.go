package number

import "math"

func reverse(x int) int {
	res := 0
	if x == math.MinInt32 {
		return 0
	}
	if x < 0 {
		x = -x
		for x > 0 {
			res = res*10 + x%10
			x = x / 10
		}
		res = -res
		if res < math.MinInt32 {
			return 0
		}
		return res
	} else {
		for x > 0 {
			res = res*10 + x%10
			x = x / 10
		}
		if res > math.MaxInt32 {
			return 0
		}
		return res
	}
}
