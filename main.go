/*
 * @Author: NyanCatda
 * @Date: 2021-12-05 22:27:13
 * @LastEditTime: 2022-06-20 13:17:31
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \MCBE-Server-Motd\main.go
 */
package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	RunPort := flag.Int("port", 8080, "指定运行端口")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 注册路由
	r = SetupRouter(r)

	fmt.Println("程序已经运行在" + strconv.Itoa(*RunPort) + "端口")
	if err := r.Run(":" + strconv.Itoa(*RunPort)); err != nil {
		fmt.Println(err)
	}
}
