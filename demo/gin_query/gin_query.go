package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/user/search", func(c *gin.Context) {
        username := c.DefaultQuery("username", "my username")
        address := c.Query("address")

        c.JSON(200, gin.H{
            "message": "pong",
            "username": username,
            "address": address,
        })
    })

    r.GET("/user/search/:username/:address", func(c *gin.Context) {
        username := c.Param("username")
        address := c.Param("address")

        // 输出json结果给调用方
        c.JSON(200, gin.H{
            "message": "pong",
            "username": username,
            "address": address,
        })
    })

    // 表单参数
    r.POST("/user/search", func(c *gin.Context) {
        //username := c.DefaultPostForm("username", "hello")
        username := c.PostForm("username")
        address := c.PostForm("address")
        c.JSON(200, gin.H{
            "message": "pong",
            "username": username,
            "address": address,
        })
    })

    r.Run() // listen and serve on 0.0.0.0:8080
}
