package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// 定义一个 map ，用于存储数值和对应的索引
	numMap := make(map[int]int)
	// 遍历数组
	for i, num := range nums {
		// 计算需要找到的目标值
		complement := target - num
		// 如果目标值在 map 中存在，说明找到了符合条件的两个数
		if j, ok := numMap[complement]; ok {
			return []int{j, i}
		}
		// 将当前数值和索引存入 map
		numMap[num] = i
	}
	// 如果没有找到符合条件的两个数，返回空切片
	return []int{}
}

func main() {

	// 示例 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := twoSum(nums1, target1)
	fmt.Println("示例 1 输出:", result1)

	// 示例 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Println("示例 2 输出:", result2)

	// 示例 3
	nums3 := []int{3, 3}
	target3 := 6
	result3 := twoSum(nums3, target3)
	fmt.Println("示例 3 输出:", result3)
}
