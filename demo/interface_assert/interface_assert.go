package main

import "fmt"

func main() {
	var a int = 100
	test(a)

	var b string = "hello"
	test(b)
}

func test(a interface{}) {
	s, ok := a.(int)
	if !ok {
		fmt.Printf("a can't convert to int\n")
		return
	}
	fmt.Println(s)

	//fmt.Printf("%T\n", a)
	//fmt.Printf("%T\n", s)
}
