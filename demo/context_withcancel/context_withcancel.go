package main

import (
    "context"
    "log"
    "os"
    "time"
)

var logg *log.Logger

func main() {
    logg = log.New(os.Stdout, "", log.Ltime)
    someHandler()
    logg.Printf("down")
}

func someHandler() {
    ctx, cancel := context.WithCancel(context.Background())
    go doStuff(ctx)
    time.Sleep(time.Second * 10)
    cancel()
}

func doStuff(ctx context.Context) {
    for {
        time.Sleep(time.Second)
        select {
        case <- ctx.Done():
            logg.Printf("done")
            return
        default:
            logg.Printf("work")
        }
    }
}
