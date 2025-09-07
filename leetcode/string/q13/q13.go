package main

import "fmt"

func romanToInt(s string) int {
	// 建立罗马数字字符到整数的映射
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	length := len(s)

	// 遍历罗马数字字符串
	for i := 0; i < length; i++ {
		// 获取当前字符对应的数值
		current := romanMap[s[i]]
		// 如果不是最后一个字符，且当前字符数值小于下一个字符数值，说明是特殊情况（如 IV、 IX 等）
		if i < length-1 && current < romanMap[s[i+1]] {
			// 用下一个字符数值减去当前字符数值，并累加到结果
			result += romanMap[s[i+1]] - current
			// 因为已经处理了下一个字符，所以 i 要加 1
			i++
		} else {
			// 普通情况，直接将当前字符数值累加到结果
			result += current
		}
	}

	return result
}

func main() {
	// 测试示例
	testCases := []struct {
		s    string
		want int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tc := range testCases {
		got := romanToInt(tc.s)
		fmt.Printf("Input: s = \"%s\", Output: %d, Expected: %d\n", tc.s, got, tc.want)
	}
}
