package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Binding from JSON
type Login struct {
    User string `form:"user" json:"user" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

func main() {
    r := gin.Default()

    // Example for binding JSON ({"user": "manu", "password": "123"})
    r.POST("/loginJSON", func(c *gin.Context) {
        var login Login
        if err := c.ShouldBindJSON(&login); err == nil {
            c.JSON(http.StatusOK, gin.H{
                "user": login.User,
                "password": login.Password,
            })
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })

    // Example for binding a HTML form (user=manu&password=123)
    r.POST("/loginForm", func(c *gin.Context) {
        var login Login
        // This will infer what binder to use depending on the content-type header.
        if err := c.ShouldBind(&login); err == nil {
            c.JSON(http.StatusOK, gin.H{
                "user": login.User,
                "password": login.Password,
            })
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })

    // Example for binding a HTML querystring (user=manu&password=123)
    r.GET("/loginForm", func(c *gin.Context) {
        var login Login
        if err := c.ShouldBind(&login); err == nil {
            c.JSON(http.StatusOK, gin.H{
                "user": login.User,
                "password": login.Password,
            })
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })

    r.Run()
}
