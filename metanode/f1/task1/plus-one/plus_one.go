package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始遍历
	for i := n - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++
		//取余判断是否有进位
		digits[i] %= 10
		// 如果当前位加1后对10取余结果不是0，说明没有进位，直接返回
		if digits[i] != 0 {
			return digits
		}
	}

	// 如果所有位都进位了，需要在最前面添加一个 1
	digits = append([]int{1}, digits...)
	return digits
}

func main() {
	// Example 1
	digits1 := []int{1, 2, 3}
	fmt.Println(plusOne(digits1)) // Output: [1, 2, 4]

	// Example 2
	digits2 := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits2)) // Output: [4, 3, 2, 2]

	// Example 3
	digits3 := []int{9}
	fmt.Println(plusOne(digits3)) // Output: [1, 0]
}
