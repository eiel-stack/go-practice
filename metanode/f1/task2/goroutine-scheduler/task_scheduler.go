package main

import (
	"fmt"
	"sync"
	"time"
)

// 任务调度器，接收任务列表并并发执行，统计执行时间
func taskScheduler(tasks []func()) {
	var wg sync.WaitGroup
	for i, task := range tasks {
		wg.Add(1)
		go func(taskIndex int, taskFunc func()) {
			defer wg.Done()
			startTime := time.Now()
			taskFunc()
			elapsedTime := time.Since(startTime)
			fmt.Printf("任务 %d 执行时间：%v \n", taskIndex+1, elapsedTime)
		}(i, task)
	}
	wg.Wait()
}

func main() {
	// 定义一些示例任务
	task1 := func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务 1 执行完毕")
	}
	task2 := func() {
		time.Sleep(2 * time.Second)
		fmt.Println("任务 2 执行完毕")
	}
	task3 := func() {
		time.Sleep(1500 * time.Millisecond)
		fmt.Println("任务 3 执行完毕")
	}

	tasks := []func(){task1, task2, task3}
	// 调度任务
	taskScheduler(tasks)
}
