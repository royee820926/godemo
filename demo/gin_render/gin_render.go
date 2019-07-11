package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    // JSON render
    // gin.H is a shortcat for map[string]interface{}
    //r.GET("/someJSON", func(c *gin.Context) {
    //    // 第一种方式，自己拼JSON
    //    c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
    //})

    // JSON render
    //r.GET("/moreJSON", func(c *gin.Context) {
    //    // you also can use a struct
    //    var msg struct {
    //        Name string `json:"user"`
    //        Message string
    //        Number int
    //    }
    //    msg.Name = "Lena"
    //    msg.Message = "hey"
    //    msg.Number = 123
    //    // Note that msg.Name becomes "user" in the JSON
    //    c.JSON(http.StatusOK, msg)
    //})

    // XML render
    //r.GET("/moreXML", func(c *gin.Context) {
    //    // you also can use a struct
    //    type MessageRecord struct {
    //        Name string
    //        Message string
    //        Number int
    //    }
    //
    //    var msg MessageRecord
    //    msg.Name = "Lena"
    //    msg.Message = "hey"
    //    msg.Number = 123
    //    c.XML(http.StatusOK, msg)
    //})

    // HTML render
    r.LoadHTMLGlob("templates/**/*")
    r.GET("posts/index", func(c *gin.Context) {
       c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
           "title": "Posts",
       })
    })
    r.GET("/users/index", func(c *gin.Context) {
       c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
           "title": "Users",
       })
    })

    // Views Test
    //r.LoadHTMLGlob("views/**/*")
    //r.GET("/msg/index", func(c *gin.Context) {
    //    c.HTML(http.StatusOK, "msg/index.html", gin.H{
    //        "title": "测试",
    //        "data":  "lists",
    //    })
    //})

    // 静态资源服务
    // 访问：localhost:8080/static/1.jpg
    r.Static("/static", "./static")

    r.Run(":8080")
}
