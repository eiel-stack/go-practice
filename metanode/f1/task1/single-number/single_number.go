package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	// 示例一
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))

	// 示例二
	nums = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))

	// 示例三
	nums = []int{1}
	fmt.Println(singleNumber(nums))
}
