package array

import "math"

func minMoves(nums []int) int {
	minN := math.MaxInt32
	sum := 0
	for _, n := range nums {
		if n < minN {
			minN = n
		}
	}
	for _, n := range nums {
		sum += n - minN
	}
	return sum
}
