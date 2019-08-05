package main

import (
    "github.com/nsqio/nsq/nsqlookupd"
)

var producer *nsqlookupd.Producer

//func main() {
//    nsqAddress := "127.0.0.1:4150"
//    err := initProducer(nsqAddress)
//    if err != nil {
//        fmt.Printf("init producer failer, err:%v\n", err)
//        return
//    }
//
//    // 读取控制台输入
//    reader := bufio.NewReader(os.Stdin)
//    for {
//        data, err := reader.ReadString('\n')
//        if err != nil {
//            fmt.Printf("read string failed, err:%v\n", err)
//            continue
//        }
//    }
//}
