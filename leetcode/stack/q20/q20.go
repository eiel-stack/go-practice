package main

import "fmt"

func isValid(s string) bool {
	// 定义一个栈，用于存储左括号
	stack := []rune{}
	// 建立右括号到左括号的映射
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 遍历字符串中的每个字符
	for _, char := range s {
		// 如果是左括号，入栈
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			// 如果栈为空，说明没有对应的左括号，返回 false
			if len(stack) == 0 {
				return false
			}
			// 取出栈顶元素
			top := stack[len(stack)-1]
			// 如果栈顶元素不是对应的左括号，返回false
			if top != bracketMap[char] {
				return false
			}
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		}
	}

	// 最后如果栈为空，说明所有括号正确匹配，返回 true，否则返回 false
	return len(stack) == 0
}

func main() {
	testCases := []struct {
		s      string
		result bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, testCase := range testCases {
		result := isValid(testCase.s)
		if result != testCase.result {
			fmt.Printf("Error: for input %s, expected %t, got %t\n", testCase.s, testCase.result, result)
		} else {
			fmt.Printf("Success: for input %s, expected %t, got %t\n", testCase.s, testCase.result, result)
		}
	}
}
