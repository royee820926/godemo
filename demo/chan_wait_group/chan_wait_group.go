package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * 一组线程都执行完成后返回
 */
func main() {
	no := 3
	var wg sync.WaitGroup
	for i:=0; i<no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(time.Second * 2)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}
