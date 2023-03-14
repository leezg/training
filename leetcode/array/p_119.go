package array

//秒！但可读性
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}

//可读性更高
//func getRow(rowIndex int) []int {
//	var pre, cur []int
//	for i := 0; i <= rowIndex; i++ {
//		cur = make([]int, i+1)
//		cur[0], cur[i] = 1, 1
//		for j := 1; j < i; j++ {
//			cur[j] = pre[j-1] + pre[j]
//		}
//		pre = cur
//	}
//	return pre
//}
