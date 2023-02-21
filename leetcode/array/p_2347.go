package array

import "sort"

const (
	flush        = "Flush"
	threeOfAKind = "Three of a Kind"
	pair         = "Pair"
	highCard     = "High Card"
)

func bestHand(ranks []int, suits []byte) string {
	sort.Ints(ranks)
	sort.Slice(suits, func(i, j int) bool {
		return suits[i] < suits[j]
	})
	ok := true
	for i := 0; i < len(suits)-1; i++ {
		if suits[i] != suits[i+1] {
			ok = false
		}
	}
	if ok {
		return flush
	}
	for i := 0; i < len(ranks)-2; i++ {
		if ranks[i] == ranks[i+1] {
			i++
			ok = true
			if ranks[i] == ranks[i+1] {
				return threeOfAKind
			}
		}
	}
	if !ok && ranks[len(ranks)-2] == ranks[len(ranks)-1] {
		ok = true
	}
	if ok {
		return pair
	}
	return highCard
}
