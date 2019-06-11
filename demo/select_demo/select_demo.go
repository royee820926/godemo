package main

import (
	"fmt"
	"time"
)

func main() {
	//testRead()

	testWrite()

	// deadlock 死锁
	//select {
	//
	//}
}

/******************** select 写 ******************/

func testWrite() {
	output1 := make(chan string, 10)
	go write(output1)

	/******************************/
	//output1 <- "hello 1"
	//output1 <- "hello 2"
	//
	//select {
	//case out1 := <- output1:
	//	fmt.Println("for recv:", out1)
	//}
	/******************************/

	for s := range output1 {
		fmt.Println("recv:", s)
		//time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write success")
		default:
			fmt.Println("channel is full")
			//break
		}
		break
	}
}

/******************** select 读 ******************/

func testRead() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	// 不带select读取管道
	//s1 := <-output1
	//fmt.Println("s1:", s1)
	//
	//s2 := <-output2
	//fmt.Println("s2:", s2)

	// select读取管道
	select {
	case s1 := <-output1:
		fmt.Println("s1:", s1)
	case s2 := <-output2:
		fmt.Println("s2:", s2)
	}
}

func server1(ch chan string) {
	time.Sleep(time.Second * 6)
	ch <- "response from server1"
}

func server2(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "response from server2"
}