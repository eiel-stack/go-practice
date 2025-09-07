package main

import (
	"fmt"
	"time"
)

type Counter struct {
	Website      string
	Start        time.Time
	PageCounters map[string]int
}

// func main() {
// 	var c Counter
// 	c.Website = "baidu.com"

// 	c.PageCounters["/"]++
// }

// func main() {
// 	var m map[int]int
// 	m[100] = 100
// }

func main() {
	var m map[int]int
	fmt.Println(m[100])
}
