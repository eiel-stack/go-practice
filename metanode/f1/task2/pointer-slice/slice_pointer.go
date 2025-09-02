package main

import "fmt"

func multiplyByTwo(slice *[]int) {
	for i := range *slice {
		(*slice)[i] *= 2
	}
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("before change,slice is %v\n", slice)
	multiplyByTwo(&slice)
	fmt.Printf("after change,slice is %v\n", slice)
}
