package main

import (
    "fmt"
    "sync"
    "time"
)

// 读写锁
var rwlock sync.RWMutex
// 互斥锁
var mutex sync.Mutex
var x int
var wg sync.WaitGroup

/**
 * 读多写少使用读写锁
 */
func main() {
    start := time.Now().UnixNano()
    wg.Add(1)
    go write()

    time.Sleep(time.Millisecond * 5)

    for i:=0; i<10; i++ {
        wg.Add(1)
        go read(i)
    }

    wg.Wait()
    end := time.Now().UnixNano()
    cost := (end - start) / 1000 / 1000
    fmt.Println("cost:", cost, "ms")
}

func write() {
    for i:=0; i<100; i++ {
        //rwlock.Lock()
        mutex.Lock()
        x = x + 1
        time.Sleep(time.Millisecond * 10)
        //rwlock.Unlock()
        mutex.Unlock()
    }

    wg.Done()
}

func read(i int) {
    for i:=0; i<100; i++ {
        //rwlock.RLock()
        mutex.Lock()
        time.Sleep(time.Millisecond)
        //rwlock.RUnlock()
        mutex.Unlock()
    }
    wg.Done()
}
