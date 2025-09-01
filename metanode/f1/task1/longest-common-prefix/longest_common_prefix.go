package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 取第一个字符串作为基准
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		prefix = prefix[:j]

		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func main() {
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs1)) // "fl"

	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs2)) // ""
}
