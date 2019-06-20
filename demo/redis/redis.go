package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
)

func main() {
    option := redis.DialPassword("foobared")
    c, err := redis.Dial("tcp", "192.168.99.221:6379", option)
    //c, err := redis.Dial("tcp", "192.168.99.221:7000,192.168.99.221:6380,192.168.99.222:6380,192.168.99.221:6381")
    if err != nil {
        fmt.Println("conn redis failed, ", err)
        return
    }
    defer c.Close()
    _, err = c.Do("Set", "abc", 102)
    if err != nil {
        fmt.Println(err)
        return
    }
    r, err := redis.Int(c.Do("Get", "abc"))
    if err != nil {
        fmt.Println("get abc failed, ", err)
        return
    }
    fmt.Println("result is ", r)
}

