package main

import (
	"fmt"
	"runtime"
	"time"
)

var i int

func main() {
	// cpu 核数
	cpu := runtime.NumCPU()
	fmt.Println("cpu:", cpu)

	// 使用2个cpu核
	runtime.GOMAXPROCS(2)

	for i := 0; i < 10; i++ {
		go calc()
	}

	time.Sleep(time.Hour)
}

func calc() {
	for {
		i++
	}
}