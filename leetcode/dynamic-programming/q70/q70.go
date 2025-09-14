package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	first, second := 1, 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

func main() {
	n1 := 2
	fmt.Printf("n = %d, 方法数：%d\n", n1, climbStairs(n1))
	n2 := 3
	fmt.Printf("n = %d, 方法数：%d\n", n2, climbStairs(n2))
	n3 := 4
	fmt.Printf("n = %d, 方法数：%d\n", n3, climbStairs(n3))
	n4 := 5
	fmt.Printf("n = %d, 方法数：%d\n", n4, climbStairs(n4))
}
