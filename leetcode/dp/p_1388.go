package dp

var ms [][]int

func maxSizeSlices(slices []int) int {
	ms = make([][]int, len(slices))
	for i, _ := range ms {
		ms[i] = make([]int, len(slices))
	}

}

func findMs(i, j int, slices []int) int {
	if i == j {
		return 0
	}

}
