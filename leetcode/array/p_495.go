package array

import "math"

func findPoisonedDuration(timeSeries []int, duration int) int {
	lastAttack := math.MinInt
	total := 0
	for _, t := range timeSeries {
		if lastAttack+duration < t {
			lastAttack = t
			total += duration
		} else {
			total += t - lastAttack
			lastAttack = t
		}
	}
	return total
}
