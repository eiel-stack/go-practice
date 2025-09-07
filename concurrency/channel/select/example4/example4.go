package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		close(ch) // 关闭通道
	}()

	select {
	case num, ok := <-ch:
		if ok {
			fmt.Printf("接收到: %d\n", num)
		} else {
			fmt.Println("通道已关闭")
		}
	}
}
