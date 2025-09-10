package main

import "fmt"

func maxArea(height []int) int {
	maxArea := 0
	i := 0
	j := len(height) - 1
	for i < j {
		// 计算当前容器的容量
		area := (j - i) * min(height[i], height[j])

		// 更新最大容量
		if area > maxArea {
			maxArea = area
		}

		// 移动较短的垂线对应的指针
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 示例1
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height)) // 49

	// 示例2
	height = []int{1, 1}
	fmt.Println(maxArea(height)) // 1
}
