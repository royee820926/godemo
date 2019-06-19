package main

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
    "time"
)

var pool *redis.Pool

func main() {
    pool = newPool("192.168.99.221:6379", "")
    for {
        time.Sleep(time.Second)
        conn := pool.Get()
        conn.Do("set", "abc", 100)

        r, err := redis.Int(conn.Do("get", "abc"))
        if err != nil {
            fmt.Printf("do failed, err:%v\n", err)
            continue
        }
        fmt.Printf("get from redis, result:%v\n", r)
    }

    //c, err := redis.Dial("tcp", "192.168.99.221:6379")
    //if err != nil {
    //    fmt.Println("conn redis failed,", err)
    //    return
    //}
    //defer c.Close()
    //_, err = c.Do("Set", "abc", 100)
    //if err != nil {
    //    fmt.Println(err)
    //    return
    //}
    //r, err := redis.Int(c.Do("Get", "abc"))
    //if err != nil {
    //    fmt.Println("get abc failed,", err)
    //    return
    //}
    //fmt.Println("get from redis", r)
}

func newPool(server, password string) *redis.Pool {
    return &redis.Pool{
        MaxIdle: 64,    // 最大限制连接数
        MaxActive: 1000, // 峰值请求连接数
        IdleTimeout: time.Second * 240, // 空闲连接超时时间
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", server)
            if err != nil {
                return nil, err
            }
            // 用户授权
            //if _, err := c.Do("AUTH", password); err != nil {
            //    c.Close()
            //    return nil, err
            //}

            return c, err
        },
        // 一分钟测试连接
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            if time.Since(t) < time.Minute {
                return nil
            }
            _, err := c.Do("ping")
            return err
        },
    }
}