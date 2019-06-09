package main

import "fmt"

func main() {
	go hello()
	fmt.Println("main thread terminate")
}

func hello() {
	fmt.Println("hello goroutine")
}
