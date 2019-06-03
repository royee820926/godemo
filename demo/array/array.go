package main

import (
	"fmt"
)

func main(){
	testMake()
}

func testMake() {
	var a []int
	a = make([]int, 5, 10)
	a[0] = 10

	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	a = append(a, 11)
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))

	for i:=0; i<8; i++ {
		a = append(a, i)
		fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	}

	a = append(a, 1000)
	fmt.Printf("扩容后的地址：a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
}