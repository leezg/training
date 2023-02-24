package main

import (
	"fmt"
	"training/leetcode/greedy"
)

func main() {
	nums := []int{1, 5, 0, 3, 5}
	fmt.Println(greedy.MinimumOperations(nums))
}
