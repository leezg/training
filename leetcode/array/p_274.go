package array

import "sort"

type BiggerFirst struct {
	nums []int
}

func (b BiggerFirst) Len() int {
	return len(b.nums)
}

func (b BiggerFirst) Less(i, j int) bool {
	return b.nums[i] > b.nums[j]
}

func (b BiggerFirst) Swap(i, j int) {
	b.nums[i], b.nums[j] = b.nums[j], b.nums[i]
}

func hIndex(citations []int) (h int) {
	sort.Sort(BiggerFirst{nums: citations})
	for i := 0; i < len(citations) && citations[i] > h; i++ {
		h++
	}
	return
}

//func hIndex(citations []int) (h int) {
//	sort.Ints(citations)
//	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
//		h++
//	}
//	return
//}
