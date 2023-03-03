package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 4, 5}
	// i = 2
	fmt.Println(append(nums[:3], nums[5:]...))
}
