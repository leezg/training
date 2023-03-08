package jzoffer

var values [][]int

func maxValue(grid [][]int) int {
	values = make([][]int, len(grid))
	for i := range values {
		values[i] = make([]int, len(grid[i]))
	}
	return dpValue(grid, 0, 0)
}

func dpValue(grid [][]int, i, j int) int {
	if i >= len(grid) {
		return 0
	}
	if j >= len(grid[0]) {
		return 0
	}
	if values[i][j] != 0 {
		return values[i][j]
	}
	vr := grid[i][j] + dpValue(grid, i, j+1)
	vd := grid[i][j] + dpValue(grid, i+1, j)
	if vr > vd {
		values[i][j] = vr
	} else {
		values[i][j] = vd
	}
	return values[i][j]
}
