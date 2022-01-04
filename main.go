/*
 * @Author: NyanCatda
 * @Date: 2021-12-05 22:27:13
 * @LastEditTime: 2022-01-03 15:46:02
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
	"time"

	"github.com/BlackBEDevelopment/MCBE-Server-Motd/MotdBEAPI"
	"github.com/BlackBEDevelopment/MCBE-Server-Motd/StatusImg"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"

	"github.com/gin-gonic/gin"
)

func main() {
	RunPort := flag.Int("port", 8080, "指定运行端口")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//初始化缓存
	memoryStore := persist.NewMemoryStore(1 * time.Minute)

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

	//不要问为什么MotdBE可以请求Java
	r.GET("/api/java", func(c *gin.Context) {
		Host := c.Query("host")

		data, err := MotdBEAPI.MotdJava(Host)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/status_img", cache.CacheByRequestURI(memoryStore, 10*time.Second), func(c *gin.Context) {
		Host := c.Query("host")

		Img := StatusImg.ServerStatusImg(Host)
		c.String(http.StatusOK, Img.String())
	})

	r.GET("/status_img/java", cache.CacheByRequestURI(memoryStore, 10*time.Second), func(c *gin.Context) {
		Host := c.Query("host")

		Img := StatusImg.ServerStatusImgJava(Host)
		c.String(http.StatusOK, Img.String())
	})

	fmt.Println("程序已经运行在" + strconv.Itoa(*RunPort) + "端口")
	if err := r.Run(":" + strconv.Itoa(*RunPort)); err != nil {
		fmt.Println(err)
	}
}
