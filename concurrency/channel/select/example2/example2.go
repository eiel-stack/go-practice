package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	select {
	case num := <-ch:
		fmt.Printf("接收到数据: %d\n", num)
	case <-time.After(2 * time.Second):
		fmt.Println("超时，未接收到数据")
	}
}
