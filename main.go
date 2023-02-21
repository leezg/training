package main

import (
	"fmt"
	"sort"
)

func main() {
	ranks := []int{1, 5, 2, 3, 4}
	suits := []byte{'a', 'b', 'd', 'c'}
	sort.Ints(ranks)
	sort.Slice(suits, func(i, j int) bool {
		return suits[i] < suits[j]
	})

	fmt.Print(ranks)
	fmt.Println(suits)
}
