package main

import (
	"fmt"
)

func main() {
	ans := []int{1, 2, 3}
	tmp := append(ans[:0], 9)

	//414
	//nums := []int{2, 2, 3, 1}
	//array.ThirdMax(nums)

	fmt.Println(ans)
	fmt.Println(tmp)
}
