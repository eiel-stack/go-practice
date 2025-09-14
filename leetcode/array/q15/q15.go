package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		// 跳过重复的第一个元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				// 找到符合条件的三元组，加入结果集
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 跳过重复的left元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 跳过重复的right元素
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	for _, triplet := range result {
		fmt.Println(triplet)
	}

	nums2 := []int{0, 1, 1}
	fmt.Println(threeSum(nums2))

	nums3 := []int{0, 0, 0}
	fmt.Println(threeSum(nums3))
}
