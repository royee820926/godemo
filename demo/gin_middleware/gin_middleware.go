package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "time"
)

func main() {
    // 空的gin引擎实例，不包含任何中间件
    //r := gin.New()

    // 包含日志处理和错误处理的gin引擎实例
    r := gin.Default()
    r.Use(StatCost())

    r.GET("/test", func(c *gin.Context) {
        example := c.MustGet("example").(string)

        // it would print: "12345"
        log.Println(example)
        c.JSON(http.StatusOK, gin.H{
            "message": "success",
        })
    })
    r.Run()
}

func StatCost() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()

        // 可以设置一下公共参数
        c.Set("example", "12345")
        // 等其他中间件先执行
        c.Next()
        // 获取耗时
        latency := time.Since(t)
        log.Printf("total cost time: %d us", latency/1000)

    }
}
