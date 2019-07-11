package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    // Simple group: v1
    v1 := r.Group("/v1")
    {
        v1.POST("/login", login)
        v1.POST("/submit", submit)
        v1.POST("/read", read)
    }

    // Simple group: v2
    v2 := r.Group("/v2")
    {
        v2.POST("/login", login)
        v2.POST("/submit", submit)
        v2.POST("/read", read)
    }

    r.Run()
}

func login(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
        "message": "success",
    })
}

func read(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
        "message": "success",
    })
}

func submit(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
        "message": "success",
    })
}
