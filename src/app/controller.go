package main

import (
    "github.com/gin-gonic/gin"
     "github.com/limech/filesearch"
     "net/http"
     "fmt"
)

func registerRoutes() *gin.Engine {
     fmt.Println("register routes")
    r := gin.Default()
     // r.LoadHTMLFiles("templates/views/index.html")



    r.GET("/", func(c *gin.Context) {
          c.File("templates/views/index.html")
          //c.HTML(http.StatusOK, "index.html", nil)
    })

    r.GET("/files/:id", func(c *gin.Context) {
        name := c.Param("id")
        fmt.Println("GET /file handler")
        filesFound, err := filesearch.Find(name)

        if err != nil {
            c.String(http.StatusNotFound, "404")
            return
        }

        c.JSON(http.StatusOK, gin.H {
           "FileInfo": filesFound,
        })

    })

    r.Static("/public","./public")

    return r
}