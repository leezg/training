package hash

import "fmt"

func getFolderNames(names []string) []string {
	nameMap := map[string]int{}
	for i, name := range names {
		if k, ok := nameMap[name]; ok {
			var tmp string
			for {
				tmp = fmt.Sprintf("%s(%d)", name, k)
				if _, ok := nameMap[tmp]; !ok {
					break
				}
				k++
			}
			names[i] = tmp
			nameMap[tmp] = 1
			nameMap[name] = k + 1
		} else {
			nameMap[name] = 1
		}
	}
	return names
}
