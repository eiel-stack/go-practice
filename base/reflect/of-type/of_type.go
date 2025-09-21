package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
}

func main() {
	st := &Student{
		Name:  "zhangsan",
		Age:   18,
		Score: 90.5,
	}
	t := reflect.TypeOf(st).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Tag)
	}
}
