package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(time.Second * 2)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(time.Second * 2)
	}
}

func write(ch chan int) {
	for i:=0; i<5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
