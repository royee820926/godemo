package main

import (
    "fmt"
    "sync"
)

var x int
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
    wg.Add(2)
    go add()
    go add()
    wg.Wait()
    fmt.Println("x:", x)
}

func add() {
    for i:=0; i<5000; i++ {
        // 互斥锁
        mutex.Lock()
        x = x + 1
        mutex.Unlock()
    }
    wg.Done()
}
