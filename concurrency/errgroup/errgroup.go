package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// 创建一个带有上下文的 ErrGroup
	g, ctx := errgroup.WithContext(context.Background())

	// 启动第一个任务：模拟一个失败的操作
	g.Go(func() error {
		select {
		case <-time.After(1 * time.Second):
			return fmt.Errorf("task 1 failed after 1 second")
		case <-ctx.Done():
			return ctx.Err() // 如果上下文被取消，返回取消错误
		}
	})

	// 启动第二个任务：模拟一个成功的操作
	g.Go(func() error {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Task 2 completed successfully")
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	// 启动第三个任务：模拟另一个成功操作，但可能被取消
	g.Go(func() error {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Task 3 completed successfully")
			return nil
		case <-ctx.Done():
			fmt.Println("Task 3 canceled due to error in another task")
			return ctx.Err()
		}
	})

	// 等待所有任务完成，并检查错误
	if err := g.Wait(); err != nil {
		fmt.Printf("Program ended with error: %v\n", err)
	} else {
		fmt.Println("All tasks completed without errors")
	}
}
