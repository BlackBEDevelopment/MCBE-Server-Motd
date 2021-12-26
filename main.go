/*
 * @Author: NyanCatda
 * @Date: 2021-12-05 22:27:13
 * @LastEditTime: 2021-12-26 22:42:44
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
	StatusImg.ServerStatusImg("nyan.xyz")

	RunPort := flag.Int("port", 8080, "指定运行端口")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.LoadHTMLGlob("fronend/dist/static/index.html")
	r.Static("/static", "./fronend/dist/static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/api", func(c *gin.Context) {
		Host := c.Query("host")

		data := MotdBEAPI.MotdBE(Host)
		c.JSON(http.StatusOK, data)
	})

	fmt.Println("程序已经运行在" + strconv.Itoa(*RunPort) + "端口")
	if err := r.Run(":" + strconv.Itoa(*RunPort)); err != nil {
		fmt.Println(err)
	}
}
