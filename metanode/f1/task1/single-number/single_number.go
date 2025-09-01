package main

import (
	"errors"
	"fmt"
)

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func singleNumber2(nums []int) int {

	tempMap := make(map[int]int)
	for _, num := range nums {
		tempMap[num]++
	}

	for key, value := range tempMap {
		if value == 1 {
			return key
		}
	}

	return -1
}

func singleNumber3(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("输入的数组不能为空")
	}

	tempMap := make(map[int]int)
	for _, num := range nums {
		tempMap[num]++
	}

	for key, value := range tempMap {
		if value == 1 {
			return key, nil
		}
	}

	return -1, errors.New("没有找到只出现一次的数字")
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

	// 示例一
	nums = []int{2, 2, 1}
	fmt.Println(singleNumber2(nums))

	// 示例二
	nums = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber2(nums))

	// 示例三
	nums = []int{1}
	fmt.Println(singleNumber2(nums))

	// 示例一
	nums = []int{2, 2, 1}
	result, err := singleNumber3(nums)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println(result)
	}

	// 示例二
	nums = []int{4, 1, 2, 1, 2}
	result, err = singleNumber3(nums)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println(result)
	}

	// 示例三
	nums = []int{1}
	result, err = singleNumber3(nums)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println(result)
	}

	// 测试空数组的情况
	nums = []int{}
	result, err = singleNumber3(nums)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println(result)
	}
}
