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
    timeoutHandler()
    logg.Printf("down")
}

func timeoutHandler() {
    //ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
    ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 5))
    //go doTimeOutStuff(ctx)
    go doStuff(ctx)

    time.Sleep(time.Second * 10)
    cancel()
}

func doStuff(ctx context.Context) {
    for {
        time.Sleep(time.Second * 1)
        select {
        case <- ctx.Done():
            logg.Printf("done")
            return
        default:
            logg.Printf("work")
        }
    }
}