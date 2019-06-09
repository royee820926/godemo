package main

import (
	"fmt"
	"time"
)

func main() {
	//管道读写测试
	//testChan()

	//生产&消费同步测试
	//testProduceAndConsume()

	//goroutine 线程同步退出测试
	//testSync()

	// 单向管道
	//testUnilateralism()

	// 管道关闭
	testClose()
}

/****************** 管道关闭 ******************/

func testClose() {
	ch := make(chan int)
	go producer(ch)

	// ok 判断管道关闭
	//for {
	//	// 判断关闭
	//	v, ok := <-ch
	//	if !ok {
	//		fmt.Println("chan is closed")
	//		break
	//	}
	//	fmt.Println("Received ", v, ok)
	//}

	// for range 判断管道关闭
	// 没有数据时，for循环会阻塞等待
	for v := range ch {
		fmt.Println("receive: ", v)
	}

}

func producer(chnl chan int) {
	for i:=0; i<10; i++ {
		chnl <- i
	}
	close(chnl)
}

/****************** 单向管道 ******************/

func testUnilateralism() {
	chnl := make(chan int)
	go sendData(chnl)
	readData(chnl)
}

func sendData(sendch chan<-int) {
	sendch <- 10
}

func readData(sendch <-chan int) {
	// 只能接收管道数据
	data := <-sendch
	fmt.Println(data)
}

/****************** goroutine 线程同步退出测试 ******************/

func testSync() {
	var exitChan chan bool
	exitChan = make(chan bool)
	go hello(exitChan)
	fmt.Println("main thread terminate")

	<-exitChan
}

func hello(c chan bool) {
	fmt.Println("hello goroutine")
	time.Sleep(time.Second * 5)

	c <- true
}

/***************** 生产&消费同步测试 *******************/

func testProduceAndConsume() {
	var c chan int
	fmt.Printf("c=%v\n", c)

	c = make(chan int)
	go produce(c)
	//go consume(c)
	time.Sleep(time.Second * 5)
}

func produce(c chan int) {
	fmt.Println("before input...")
	c <- 1000
	fmt.Println("produce finished")
}

func consume(c chan int) {
	time.Sleep(time.Second * 2)
	data := <- c
	fmt.Println(data)
}

/****************** 管道读写测试 ******************/

func testChan() {
	var c chan int
	fmt.Printf("c=%v\n", c)

	c = make(chan int, 10)
	fmt.Printf("c=%v\n", c)

	//入管道
	c <- 100

	// 出管道
	data := <- c
	fmt.Printf("data=%v\n", data)
}