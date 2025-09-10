package main

import "fmt"

func moveZeroes(nums []int) {

	// 用于记录非零元素应该放置的位置
	j := 0
	for i := 0; i < len(nums); i++ {
		// 如果当前元素不为零
		if nums[i] != 0 {
			// 将非零元素放到 j 位置
			nums[j] = nums[i]
			// 非零元素放置位置后移
			j++
		}
	}

	// 将 j 位置及之后的元素都置为 0
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)

}
