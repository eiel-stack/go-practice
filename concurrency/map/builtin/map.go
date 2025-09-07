package main

import (
	"fmt"
	"reflect"
)

func usage1() {
	var m = make(map[string]int)
	m["a"] = 0
	fmt.Printf("a=%d; b=%d\n", m["a"], m["b"])

	av, aexisted := m["a"]
	bv, bexisted := m["b"]
	fmt.Printf("a=%d, existed: %t; b=%d, existed: %t\n", av, aexisted, bv, bexisted)
}

func main() {
	usage1()

	m := make(map[string]interface{})
	m["name"] = "wuqq"
	m["age"] = "test1"
	m["ag2"] = "test2"
	m["age3"] = "test3"

	for item := range m {
		fmt.Println("item:", reflect.ValueOf(item))
	}

}
