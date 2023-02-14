package array

func longestWPI(hours []int) int {
	ans := 0
	for i := 0; i < len(hours); i++ {
		count := 0
		for j := i; j < len(hours); j++ {
			if hours[j] > 8 {
				count++
			}
			if count*2 > j-i+1 {
				if ans < j-i+1 {
					ans = j - i + 1
				}
			}
		}
	}
	return ans
}

//func longestWPI(hours []int) int {
//	ans := 0
//	var stk []int
//	s := make([]int, len(hours)+1)
//	stk = append(stk, 0)
//	for i, h := range hours {
//		if h > 8 {
//			s[i+1] = s[i] + 1
//		} else {
//			s[i+1] = s[i] - 1
//		}
//		if s[stk[len(stk)-1]] > s[i+1] {
//			stk = append(stk, i+1)
//		}
//	}
//	for i := len(hours); i >= 1; i-- {
//		for len(stk) > 0 && s[stk[len(stk)-1]] < s[i] {
//			if i-stk[len(stk)-1] > ans {
//				ans = stk[len(stk)-1]
//			}
//			stk = stk[:len(stk)-1]
//		}
//	}
//	return ans
//}
