package main

import (
    "context"
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

const TraceId = "trace_id"

func main() {
    http.HandleFunc("/", lazyHandler)
    http.ListenAndServe(":8000", nil)
}

func lazyHandler(w http.ResponseWriter, req *http.Request) {
    ctx := context.WithValue(context.Background(), TraceId, rand.Int63())
    a(ctx)

    ranNum := rand.Intn(2)
    if ranNum == 0 {
        time.Sleep(time.Second * 6)
        fmt.Fprintf(w, "slow response, %d\n", ranNum)
        fmt.Printf("slow response, %d\n", ranNum)
        return
    }

    fmt.Fprintf(w, "quick response, %d\n", ranNum)
    fmt.Printf("quick response, %d\n", ranNum)
    return
}

func a(ctx context.Context) {
    TraceId := ctx.Value(TraceId)
    fmt.Println("trace_id:%v, process of a\n", TraceId)
    b(ctx)
}

func b(ctx context.Context) {
    TraceId := ctx.Value(TraceId)
    fmt.Println("trace_id:%v, process of b\n", TraceId)
    c(ctx)
}

func c(ctx context.Context) {
    TraceId := ctx.Value(TraceId)
    fmt.Println("trace_id:%v, process of c\n", TraceId)
    d(ctx)
}

func d(ctx context.Context) {
    TraceId := ctx.Value(TraceId)
    fmt.Println("trace_id:%v, process of d\n", TraceId)
}