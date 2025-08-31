package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	// 负数不是回文数
	if x < 0 {
		return false
	}
	// 将整数转为字符串
	s := strconv.Itoa(x)
	// 双指针判断是否为回文
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindrome2(x int) bool {
	// 负数不是回文数
	if x < 0 {
		return false
	}
	// 保存原始数字，用于最后比较
	original := x
	// 反转后的数字
	reversed := 0

	for x > 0 {
		// 取出最后一位
		lastDigit := x % 10
		// 构建反转后的数字
		reversed = reversed*10 + lastDigit
		// 去掉最后一位
		x /= 10
	}

	// 比较原始数字和反转后的数字
	return original == reversed
}

func main() {
	x := 121
	fmt.Printf("%d is palindrome: %t\n", x, isPalindrome(x))
	x = -121
	fmt.Printf("%d is palindrome: %t\n", x, isPalindrome(x))
	x = 1221
	fmt.Printf("%d is palindrome: %t\n", x, isPalindrome(x))
	x = 0
	fmt.Printf("%d is palindrome: %t\n", x, isPalindrome(x))

	x = 121
	fmt.Printf("%d is palindrome2: %t\n", x, isPalindrome(x))
	x = -121
	fmt.Printf("%d is palindrome2: %t\n", x, isPalindrome(x))
	x = 1221
	fmt.Printf("%d is palindrome2: %t\n", x, isPalindrome(x))
	x = 0
	fmt.Printf("%d is palindrome2: %t\n", x, isPalindrome(x))
}
