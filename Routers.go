/*
 * @Author: NyanCatda
 * @Date: 2022-06-20 13:12:12
 * @LastEditTime: 2022-06-28 17:01:33
 * @LastEditors: NyanCatda
 * @Description: 路由注册
 * @FilePath: \MCBE-Server-Motd\Routers.go
 */
package main

import (
	"io/ioutil"
	"net/http"
	"path"
	"time"

	"github.com/BlackBEDevelopment/MCBE-Server-Motd/MotdBEAPI"
	"github.com/BlackBEDevelopment/MCBE-Server-Motd/StatusImg"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/nyancatda/AyaLog"
)

/**
 * @description: 路由注册
 * @param {*gin.Engine} r gin引擎
 * @return {*gin.Engine} gin引擎
 */
func SetupRouter(r *gin.Engine) *gin.Engine {
	// 500错误处理
	r.Use(ServerError)

	// 注册静态资源
	r.Static("/static", "fronend/dist/static")
	r.StaticFile("/favicon.ico", "fronend/dist/static/favicon.ico")

	// 注册前端页面
	FrontendRouter(r)

	// 基岩版查询API
	r.GET("/api", func(c *gin.Context) {
		Host := c.Query("host")

		data, err := MotdBEAPI.MotdBE(Host)
		if err != nil {
			AyaLog.DeBug("Request", err)
		}
		c.JSON(http.StatusOK, data)
	})

	// Java版查询API
	// 不要问为什么MotdBE可以请求Java
	r.GET("/api/java", func(c *gin.Context) {
		Host := c.Query("host")

		data, err := MotdBEAPI.MotdJava(Host)
		if err != nil {
			AyaLog.DeBug("Request", err)
		}
		c.JSON(http.StatusOK, data)
	})

	//初始化缓存
	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	// 基岩版服务器状态图片
	r.GET("/status_img", cache.CacheByRequestURI(memoryStore, 10*time.Second), func(c *gin.Context) {
		Host := c.Query("host")

		Img, err := StatusImg.ServerStatusImg(Host)
		if err != nil {
			AyaLog.Error("Request", err)
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		c.String(http.StatusOK, Img.String())
	})

	// Java版服务器状态图片
	r.GET("/status_img/java", cache.CacheByRequestURI(memoryStore, 10*time.Second), func(c *gin.Context) {
		Host := c.Query("host")

		Img, err := StatusImg.ServerStatusImgJava(Host)
		if err != nil {
			AyaLog.Error("Request", err)
			c.String(http.StatusInternalServerError, "Server Error")
			return
		}
		c.String(http.StatusOK, Img.String())
	})

	return r
}

/**
 * @description: 500错误处理
 * @param {*gin.Context} c
 * @return {*}
 */
func ServerError(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// 打印错误堆栈信息
			AyaLog.Error("Request", r.(error))

			// 500返回
			c.String(http.StatusInternalServerError, "Server Error")
			c.Abort()
			return
		}
	}()
	c.Next()
}

/**
 * @description: 前端文件路由注册
 * @param {*gin.Engine} r 路由引擎
 * @return {*gin.Engine} 路由引擎
 */
func FrontendRouter(r *gin.Engine) *gin.Engine {
	// 遍历目录下的所有HTML文件
	HTMLFiles, err := ioutil.ReadDir("fronend/dist/static")
	if err != nil {
		AyaLog.Error("System", err)
	}
	// 注册index.html
	r.StaticFile("/", "fronend/dist/static/index.html")
	// 循环注册其他HTML文件
	for _, File := range HTMLFiles {
		// 判断是否为HTML文件
		FileName := File.Name()
		if path.Ext(FileName) == ".html" {
			r.StaticFile(FileName, "fronend/dist/static/"+FileName)
		}
	}

	return r
}
