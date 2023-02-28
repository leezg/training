package hash

import "sort"

type ByValue [][]int

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i][0] < a[j][0] }

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	var res [][]int
	sort.Sort(ByValue(items1))
	sort.Sort(ByValue(items2))
	i, j := 0, 0
	for i < len(items1) && j < len(items2) {
		if items1[i][0] < items2[j][0] {
			res = append(res, items1[i])
			i++
		} else if items1[i][0] > items2[j][0] {
			res = append(res, items2[j])
			j++
		} else {
			res = append(res, []int{items1[i][0], items1[i][1] + items2[j][1]})
			i++
			j++
		}
	}
	for i < len(items1) {
		res = append(res, items1[i])
		i++
	}
	for j < len(items2) {
		res = append(res, items2[j])
		j++
	}
	return res
}
