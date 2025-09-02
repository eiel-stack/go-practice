package main

import "fmt"

func addTen(num *int) {
	*num += 10
}

func main() {
	num := 5
	fmt.Printf("before change,num is %d\n", num)
	addTen(&num)
	fmt.Printf("after change,num is %d\n", num)
}
