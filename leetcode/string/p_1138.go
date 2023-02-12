package string

func AlphabetBoardPath(target string) string {
	return alphabetBoardPath(target)
}

func alphabetBoardPath(target string) string {
	r, c := 0, 0
	ans := ""
	for _, a := range target {
		nextR, nextC := getCoor(a)
	ForLoop:
		for r < nextR {
			if r == 4 && c != 0 {
				break ForLoop
			}
			ans += "D"
			r++
		}
		for r > nextR {
			ans += "U"
			r--
		}
		for c < nextC {
			ans += "R"
			c++
		}
		for c > nextC {
			ans += "L"
			c--
		}
		if r < nextR {
			ans += "D"
			r++
		}
		ans += "!"
	}
	return ans
}

func getCoor(alpha int32) (int, int) {
	return int((alpha - 97) / 5), int((alpha - 97) % 5)
}
