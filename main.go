package main

import (
	"fmt"
	"training/leetcode/array"
)

func main() {
	fmt.Scan()
	nums := []int{0, 0, 1}
	array.MoveZeroes(nums)
	fmt.Println(nums)
}
