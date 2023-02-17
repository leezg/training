package main

import (
	"fmt"
	"training/leetcode/array"
)

func main() {
	grid := [][]int{{1, 1, 0, 0}}

	fmt.Println(array.Largest1BorderedSquare(grid))
}
