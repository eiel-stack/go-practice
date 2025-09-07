package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		ch1 <- 42
	}()

	go func() {
		ch2 <- "hello"
	}()

	select {
	case num := <-ch1:
		fmt.Printf("从 ch1 接收到: %d\n", num)
	case str := <-ch2:
		fmt.Printf("从 ch2 接收到: %s\n", str)
	}
}
