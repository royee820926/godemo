package main

import "fmt"

func main() {
	interface_demo()
}

/**
 * 接口interface测试
 */
func interface_demo()  {
	var a interface{}
	var b int = 100
	a = b
	fmt.Printf("%T %v\n", a, a)

	var c string = "hello"
	a = c
	fmt.Printf("%T %v\n", a, a)

	var d map[string]int = make(map[string]int, 100)
	d["abc"] = 1000
	d["eke"] = 30
	a = d
	fmt.Printf("%T %v\n", a, a)
}
