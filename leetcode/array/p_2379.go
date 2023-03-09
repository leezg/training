package array

func MinimumRecolors(blocks string, k int) int {
	return minimumRecolors(blocks, k)
}

func minimumRecolors(blocks string, k int) int {
	res := 0
	i, j := 0, 0

	for ; j < k; j++ {
		if blocks[j] == 'W' {
			res++
		}
	}
	//j--
	tmp := res
	for j < len(blocks) {
		if blocks[j] == 'W' && blocks[i] == 'B' {
			tmp++
		} else if blocks[j] == 'B' && blocks[i] == 'W' {
			tmp--
		}
		if tmp < res {
			res = tmp
		}
		i++
		j++
	}
	return res
}
