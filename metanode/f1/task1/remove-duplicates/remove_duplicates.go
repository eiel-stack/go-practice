package main

import "fmt"

func removeDuplicates(nums []int) int {
	// 如果数组长度为 0，则直接返回 0
	if len(nums) == 0 {
		return 0
	}

	// 慢指针，指向当前不重复元素应该放置的位置
	i := 0
	// 快指针，用于遍历数组
	for j := 1; j < len(nums); j++ {
		// 如果快指针指向的元素与慢指针指向的元素不同
		if nums[i] != nums[j] {
			//慢指针向后移动一位
			i++
			// 将快指针指向的不重复元素放到慢指针指向的位置
			nums[i] = nums[j]
		}
	}

	// 慢指针 + 1 就是不重复元素的个数
	return i + 1
}

func main() {
	// Example 1
	nums1 := []int{1, 1, 2}
	length1 := removeDuplicates(nums1)
	fmt.Println("Example 1:", length1)
	fmt.Println(nums1[:length1]) // Output: [1 2]

	// Example 2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length2 := removeDuplicates(nums2)
	fmt.Println("Example 2:", length2)
	fmt.Println(nums2[:length2]) // Output: [0 1 2 3 4]
}
