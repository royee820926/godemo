package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

func main() {
    r := gin.Default()

    // 设置上传文件缓冲区大小，默认32 MiB
    r.MaxMultipartMemory = 8 << 20 // 8 MiB

    // 单个文件上传
    r.POST("/upload", func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "message": err.Error(),
            })
            return
        }

        log.Println(file.Filename)
        // 保存的目标文件
        dst := fmt.Sprintf("c:/tmp/%s", file.Filename)
        // upload the file to specific dst.
        c.SaveUploadedFile(file, dst)
        c.JSON(http.StatusOK, gin.H{
            "message": fmt.Sprintf("'%s' uploaded!", file.Filename),
        })
    })

    // 多个文件上传
    r.POST("/multi_upload", func(c *gin.Context) {
        // Multipart form
        form, _ := c.MultipartForm()
        files := form.File["file"]
        for index, file := range files {
            log.Println(file.Filename)
            dst := fmt.Sprintf("c:/tmp/%s_%d", file.Filename, index)
            c.SaveUploadedFile(file, dst)
        }
        c.JSON(http.StatusOK, gin.H{
            "message": fmt.Sprintf("%d files uploaded!", len(files)),
        })
    })

    r.Run()
}
