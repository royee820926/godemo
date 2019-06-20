package main

import (
    "github.com/gitstliu/go-redis-cluster"
    "time"
)

func main() {
    cluster, _ := redis.NewCluster(
        &redis.Options{
            StartNodes: []string{
                "192.168.99.221:7000",
                "192.168.99.221:6380",
                "192.168.99.222:6380",
                "192.168.99.221:6381",

                "192.168.99.222:6381",
                "192.168.99.221:7001",
                "192.168.99.223:6381",
                "192.168.99.223:6380",
            },
            ConnTimeout: time.Millisecond * 50,
            ReadTimeout: time.Millisecond * 50,
            WriteTimeout:time.Millisecond * 50,
            KeepAlive: 16,
            AliveTime: time.Second * 60,
        })
    //if err != nil {
    //    fmt.Printf("new cluster failed, err:%v\n", err)
    //    return
    //}

    cluster.Do("AUTH", "foobared")
    cluster.Do("SET", "foo", "bar")
}
