// 超时测试服务端：context_server_timeout
// 超时测试客户端：context_client_timeout

package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", lazyHandler)
    http.ListenAndServe(":8000", nil)
}

func lazyHandler(w http.ResponseWriter, req *http.Request) {
    ranNum := rand.Intn(2)
    if ranNum == 0 {
        time.Sleep(time.Second * 6)
        fmt.Fprintf(w, "slow response, %d\n", ranNum)
        fmt.Printf("slow response, %d\n", ranNum)
        return
    }
}