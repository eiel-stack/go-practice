package main

import "fmt"

func isValid(s string) bool {
	// 定义一个栈，用于存储左括号
	var stack []rune
	// 定义括号的映射关系，右括号对应左括号
	bracketMap := map[rune]rune{
		')': '(', ']': '[', '}': '{',
	}

	for _, char := range s {
		// 如果是左括号，入栈
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			// 如果栈为空，说明没有对应的左括号，返回 false
			if len(stack) == 0 {
				return false
			}

			// 获取栈顶元素
			top := stack[len(stack)-1]
			// 如果右括号对应的左括号与栈顶元素不匹配，返回 false
			if bracketMap[char] != top {
				return false
			}
			// 匹配成功，栈顶元素出栈
			stack = stack[:len(stack)-1]
		}
	}

	// 最后如果栈为空，说明所有括号都匹配
	return len(stack) == 0
}

func main() {

	// TEST CASE
	testCases := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, tc := range testCases {
		result := isValid(tc.s)
		fmt.Printf("输入: %s, 期望输出: %t, 实际输出: %t\n", tc.s, tc.expected, result)
	}
}
