package array

func generate(numRows int) (res [][]int) {
	res = make([][]int, numRows)
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		if i == 1 {
			res[i] = []int{1, 1}
			continue
		}
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return
}
