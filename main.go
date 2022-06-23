/*
 * @Author: NyanCatda
 * @Date: 2021-12-05 22:27:13
 * @LastEditTime: 2022-06-23 23:39:57
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \MCBE-Server-Motd\main.go
 */
package main

import (
	"flag"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/AyaLog/ModLog/GinLog"
	"github.com/nyancatda/AyaLog/TimedTask"
)

func main() {
	RunPort := flag.Int("port", 8080, "指定运行端口")
	DeBug := flag.Bool("debug", false, "是否开启调试模式")
	flag.Parse()

	// 设置日志参数
	if !*DeBug {
		AyaLog.LogLevel = AyaLog.INFO
	}

	// 启用日志压缩与清理任务
	go TimedTask.Start()

	// 关闭Gin默认的日志输出
	gin.DefaultWriter = ioutil.Discard

	// 初始化Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 注册日志中间件
	r.Use((GinLog.GinLog()))

	// 注册路由
	SetupRouter(r)

	AyaLog.Info("System", "程序启动完成！正在监听端口："+strconv.Itoa(*RunPort))
	if err := r.Run(":" + strconv.Itoa(*RunPort)); err != nil {
		AyaLog.Error("System", err)
	}
}
