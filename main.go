/*
 * @Author: NyanCatda
 * @Date: 2021-12-05 22:27:13
 * @LastEditTime: 2021-12-28 12:45:44
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \MotdBE\main.go
 */
package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BlackBEDevelopment/MCBE-Server-Motd/MotdBEAPI"
	"github.com/BlackBEDevelopment/MCBE-Server-Motd/StatusImg"

	"github.com/gin-gonic/gin"
)

func main() {
	RunPort := flag.Int("port", 8080, "指定运行端口")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.LoadHTMLGlob("fronend/dist/static/**.html")
	r.Static("/static", "./fronend/dist/static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/iframe", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iframe.html", gin.H{})
	})
	r.GET("/iframe.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iframe.html", gin.H{})
	})

	r.GET("/api", func(c *gin.Context) {
		Host := c.Query("host")

		data, err := MotdBEAPI.MotdBE(Host)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/api/java", func(c *gin.Context) {
		Host := c.Query("host")

		data, err := MotdBEAPI.MotdJava(Host)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/status_img", func(c *gin.Context) {
		Host := c.Query("host")

		Img := StatusImg.ServerStatusImg(Host)
		c.String(http.StatusOK, Img.String())
	})

	fmt.Println("程序已经运行在" + strconv.Itoa(*RunPort) + "端口")
	if err := r.Run(":" + strconv.Itoa(*RunPort)); err != nil {
		fmt.Println(err)
	}
}
