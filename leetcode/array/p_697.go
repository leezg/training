package array

func FindShortestSubArray(nums []int) int {
	return findShortestSubArray(nums)
}

type table struct {
	length int
	begin  int
	end    int
}

func findShortestSubArray(nums []int) int {
	degree := map[int]table{}
	for i, n := range nums {
		if _, ok := degree[n]; ok {
			degree[n] = table{
				length: degree[n].length + 1,
				begin:  degree[n].begin,
				end:    i,
			}
		} else {
			degree[n] = table{
				length: 1,
				begin:  i,
				end:    i,
			}
		}
	}
	ans := 0
	maxCnt := 0
	for _, e := range degree {
		if e.length > maxCnt {
			maxCnt, ans = e.length, e.end-e.begin+1
		} else if e.length == maxCnt {
			ans = min(ans, e.end-e.begin+1)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
