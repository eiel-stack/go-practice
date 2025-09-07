package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 10 // 填充一个值

	// 非阻塞发送
	select {
	case ch <- 20:
		fmt.Println("发送成功")
	default:
		fmt.Println("通道已满，发送失败")
	}

	// 非阻塞接收
	select {
	case num := <-ch:
		fmt.Printf("接收到: %d\n", num)
	default:
		fmt.Println("通道为空，接收失败")
	}

	// 非阻塞发送
	select {
	case ch <- 20:
		fmt.Println("发送成功")
	default:
		fmt.Println("通道已满，发送失败")
	}

}
