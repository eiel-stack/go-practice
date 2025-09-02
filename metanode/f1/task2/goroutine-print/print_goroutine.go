package main

import (
	"fmt"
	"sync"
)

// 打印奇数的协程函数
func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Odd:", i)
	}
}

// 打印偶数的协程函数
func printEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println("Even:", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// 启动打印奇数的协程
	go printOdd(&wg)
	// 启动打印偶数的协程
	go printEven(&wg)
	// 等待两个协程执行完毕
	wg.Wait()
}
