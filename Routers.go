/*
 * @Author: NyanCatda
 * @Date: 2022-06-20 13:12:12
 * @LastEditTime: 2022-06-20 13:17:02
 * @LastEditors: NyanCatda
 * @Description: 路由注册
 * @FilePath: \MCBE-Server-Motd\Routers.go
 */
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/BlackBEDevelopment/MCBE-Server-Motd/MotdBEAPI"
	"github.com/BlackBEDevelopment/MCBE-Server-Motd/StatusImg"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

/**
 * @description: 路由注册
 * @param {*gin.Engine} r gin引擎
 * @return {*gin.Engine} gin引擎
 */
func SetupRouter(r *gin.Engine) *gin.Engine {
	// 注册静态资源
	r.Static("/static", "./fronend/dist/static")

	// 注册HTML资源
	r.LoadHTMLGlob("fronend/dist/static/**.html")

	// 主页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	// iframe
	r.GET("/iframe", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iframe.html", gin.H{})
	})
	// iframe
	r.GET("/iframe.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iframe.html", gin.H{})
	})

	// 基岩版查询API
	r.GET("/api", func(c *gin.Context) {
		Host := c.Query("host")

		data, err := MotdBEAPI.MotdBE(Host)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, data)
	})

	// Java版查询API
	// 不要问为什么MotdBE可以请求Java
	r.GET("/api/java", func(c *gin.Context) {
		Host := c.Query("host")

		data, err := MotdBEAPI.MotdJava(Host)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, data)
	})

	//初始化缓存
	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	// 基岩版服务器状态图片
	r.GET("/status_img", cache.CacheByRequestURI(memoryStore, 10*time.Second), func(c *gin.Context) {
		Host := c.Query("host")

		Img := StatusImg.ServerStatusImg(Host)
		c.String(http.StatusOK, Img.String())
	})

	// Java版服务器状态图片
	r.GET("/status_img/java", cache.CacheByRequestURI(memoryStore, 10*time.Second), func(c *gin.Context) {
		Host := c.Query("host")

		Img := StatusImg.ServerStatusImgJava(Host)
		c.String(http.StatusOK, Img.String())
	})

	return r
}
