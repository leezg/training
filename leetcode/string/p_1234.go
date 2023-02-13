package string

import "math"

func BalancedString(s string) int {
	return balancedString(s)
}

func balancedString(s string) int {
	var total map[int32]int
	var over map[int32]int
	var count map[int32]int
	total, over, count = make(map[int32]int), make(map[int32]int), make(map[int32]int)
	total['Q'], total['W'], total['E'], total['R'] = 0, 0, 0, 0
	a := len(s) / 4
	for _, c := range s {
		total[c]++
	}
	over['Q'], over['W'], over['E'], over['R'] = total['Q']-a, total['W']-a, total['E']-a, total['R']-a
	if over['Q'] == 0 && over['W'] == 0 && over['E'] == 0 && over['R'] == 0 {
		return 0
	}
	res := math.MaxInt
	l, r := 0, 0
	for ; r < len(s); r++ {
		count[int32(s[r])]++
		if count['Q'] >= over['Q'] && count['W'] >= over['W'] && count['E'] >= over['E'] && count['R'] >= over['R'] {
			if l == r {
				return 1
			}
		Loop1:
			for ; l < r; l++ {
				count[int32(s[l])]--
				if count['Q'] < over['Q'] || count['W'] < over['W'] || count['E'] < over['E'] || count['R'] < over['R'] {
					if res > r-l+1 {
						res = r - l + 1
					}
					l++
					break Loop1
				}
			}
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}
