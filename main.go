package main

import (
	"blackbe.xyz/MotdBE/MotdBEAPI"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.LoadHTMLGlob("assets/*")
	r.Static("/assets", "./assets")
	r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{})
    })

	r.GET("/api", func(c *gin.Context) {
		Host := c.Query("host")

		data := MotdBEAPI.MotdBE(Host)
		c.JSON(http.StatusOK, data)
	})

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
