/*
 * @Author: NyanCatda
 * @Date: 2021-12-26 21:23:59
 * @LastEditTime: 2021-12-26 22:47:20
 * @LastEditors: NyanCatda
 * @Description: 服务器状态图片生成
 * @FilePath: \MotdBE\StatusImg\StatusImg.go
 */
package StatusImg

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/BlackBEDevelopment/MCBE-Server-Motd/MotdBEAPI"
	"github.com/golang/freetype"
)

func ServerStatusImg(Host string) {
	imgfile, _ := os.Create("test.png")
	defer imgfile.Close()

	//读取背景图片
	backgroundFile, err := os.Open("StatusImg/background.png")
	if err != nil {
		panic(err)
	}
	backgroundImg, err := png.Decode(backgroundFile)
	if err != nil {
		panic(err)
	}

	//转换类型
	img := backgroundImg.(*image.NRGBA)

	//读取字体数据
	fontBytes, err := ioutil.ReadFile("StatusImg/SourceHanSansCN-VF.ttf")
	if err != nil {
		log.Println(err)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println("load front fail", err)
	}

	f := freetype.NewContext()
	//设置分辨率
	f.SetDPI(72)
	//设置字体
	f.SetFont(font)
	//设置尺寸
	f.SetFontSize(40)
	f.SetClip(img.Bounds())
	//设置输出的图片
	f.SetDst(img)
	//设置字体颜色(白色)
	f.SetSrc(image.NewUniform(color.RGBA{255, 255, 255, 255}))

	//获取服务器信息
	ServerData := MotdBEAPI.MotdBE(Host)

	//设置字体的位置
	pt := freetype.Pt(10, 40+int(f.PointToFixed(26))>>8)
	f.DrawString("MOTD: "+ServerData.Motd, pt)
	pt = freetype.Pt(10, 90+int(f.PointToFixed(26))>>8)
	f.DrawString("协议版本: "+ServerData.Agreement, pt)
	pt = freetype.Pt(10, 140+int(f.PointToFixed(26))>>8)
	f.DrawString("游戏版本: "+ServerData.Version, pt)
	pt = freetype.Pt(10, 190+int(f.PointToFixed(26))>>8)
	f.DrawString("在线人数: ", pt)
	pt = freetype.Pt(10, 240+int(f.PointToFixed(26))>>8)
	f.DrawString("存档名字: "+ServerData.LevelName, pt)
	pt = freetype.Pt(10, 290+int(f.PointToFixed(26))>>8)
	f.DrawString("游戏模式: "+ServerData.GameMode, pt)
	pt = freetype.Pt(10, 340+int(f.PointToFixed(26))>>8)
	f.DrawString("连接延迟: "+strconv.FormatInt(ServerData.Delay, 10), pt)

	//以png 格式写入文件
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}
