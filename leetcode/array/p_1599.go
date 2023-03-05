package array

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	res := 0
	total := 0
	i := 0
	times := -1
	remain := 0

	for ; i < len(customers); i++ {
		remain += customers[i]
		if remain > 4 {
			total += 4*boardingCost - runningCost
			remain -= 4
		} else {
			total += remain*boardingCost - runningCost
			remain = 0
		}
		if total > res {
			res = total
			times = i + 1
		}
	}
	for remain > 4 {
		total += 4*boardingCost - runningCost
		remain -= 4
		i++
		if total > res {
			res = total
			times = i
		}
	}
	if remain > 0 {
		total += remain*boardingCost - runningCost
		remain = 0
		i++
		if total > res {
			res = total
			times = i
		}
	}
	return times
}
