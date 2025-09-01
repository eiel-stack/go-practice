package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 如果区间数组长度小于等于1，直接返回原数组
	if len(intervals) <= 1 {
		return intervals
	}

	// 按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 结果集，存储合并后的区间
	result := [][]int{intervals[0]}

	// 遍历排序后的区间数组
	for i := 1; i < len(intervals); i++ {
		// 获取结果集中最后一个区间
		last := result[len(result)-1]
		// 如果当前区间的起始位置小于等于结果集中最后一个区间的结束位置，说明有重叠
		if intervals[i][0] <= last[1] {
			// 合并后的区间位置取两者的较大者
			if last[1] < intervals[i][1] {
				last[1] = intervals[i][1]
			}
		} else {
			// 没有重叠，将当前区间加入结果集
			result = append(result, intervals[i])
		}
	}

	return result
}

func main() {
	// Example 1
	intervals := [][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))

	// Example 2
	intervals = [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals))

	// 示例 3
	intervals3 := [][]int{{4, 7}, {1, 4}}
	result3 := merge(intervals3)
	fmt.Println("示例 3 输出:", result3)
}
