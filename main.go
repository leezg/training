package main

import (
	"fmt"
	"training/leetcode/string"
)

func main() {
	s := "WQWRQQQW"
	ans := string.BalancedString(s)

	fmt.Println(ans)
}
