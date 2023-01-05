package main

import (
	"fmt"
	"training/leetcode/array"
)

func main() {
	ans := []int{1, 2, 3}
	tmp := append(ans[:0], 9)

	//414
	nums := []int{1, 2, 2, 4}
	array.FindErrorNums(nums)

	fmt.Println(ans)
	fmt.Println(tmp)
}
