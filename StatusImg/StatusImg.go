/*
 * @Author: NyanCatda
 * @Date: 2021-12-26 21:23:59
 * @LastEditTime: 2022-06-20 13:28:28
 * @LastEditors: NyanCatda
 * @Description: 服务器状态图片生成
 * @FilePath: \MCBE-Server-Motd\StatusImg\StatusImg.go
 */
package StatusImg

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/BlackBEDevelopment/MCBE-Server-Motd/MotdBEAPI"
	"github.com/golang/freetype"
)

/**
 * @description: 服务器状态图片生成
 * @param {string} Host 服务器地址
 * @return {*bytes.Buffer} 图片Buffer
 */
func ServerStatusImg(Host string) (*bytes.Buffer, error) {
	//获取服务器信息
	ServerData, err := MotdBEAPI.MotdBE(Host)
	if err != nil {
		return nil, err
	}
	if ServerData.Status == "offline" {
		offlineImgFile, err := os.Open("StatusImg/background.png")
		if err != nil {
			return nil, err
		}
		offlineImg, err := png.Decode(offlineImgFile)
		if err != nil {
			return nil, err
		}
		//将图片写入Buffer
		Buff := bytes.NewBuffer(nil)
		err = png.Encode(Buff, offlineImg)
		if err != nil {
			return nil, err
		}
		return Buff, nil
	}

	//读取背景图片
	backgroundFile, err := os.Open("StatusImg/background.png")
	if err != nil {
		return nil, err
	}
	backgroundImg, err := png.Decode(backgroundFile)
	if err != nil {
		return nil, err
	}

	//转换类型
	img := backgroundImg.(*image.NRGBA)

	//读取字体数据
	fontBytes, err := ioutil.ReadFile("StatusImg/unifont-12.1.04.ttf")
	if err != nil {
		return nil, err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	//设置标题字体
	f := freetype.NewContext()
	//设置分辨率
	f.SetDPI(72)
	//设置字体
	f.SetFont(font)
	//设置尺寸
	f.SetFontSize(30)
	f.SetClip(img.Bounds())
	//设置输出的图片
	f.SetDst(img)
	//设置字体颜色(白色)
	f.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 255}))
	pt := freetype.Pt(20, 30+int(f.PointToFixed(26))>>8)
	f.DrawString("MOTD: "+RemoveColorCode(ServerData.Motd), pt)

	//设置内容字体
	f = freetype.NewContext()
	//设置分辨率
	f.SetDPI(72)
	//设置字体
	f.SetFont(font)
	//设置尺寸
	f.SetFontSize(30)
	f.SetClip(img.Bounds())
	//设置输出的图片
	f.SetDst(img)
	//设置字体颜色(白色)
	f.SetSrc(image.NewUniform(color.RGBA{255, 255, 255, 255}))
	pt = freetype.Pt(20, 75+int(f.PointToFixed(26))>>8)
	f.DrawString("协议版本: "+strconv.Itoa(ServerData.Agreement), pt)
	pt = freetype.Pt(20, 125+int(f.PointToFixed(26))>>8)
	f.DrawString("游戏版本: "+ServerData.Version, pt)
	pt = freetype.Pt(20, 175+int(f.PointToFixed(26))>>8)
	f.DrawString("在线人数: "+strconv.Itoa(ServerData.Online)+"/"+strconv.Itoa(ServerData.Max), pt)
	pt = freetype.Pt(20, 225+int(f.PointToFixed(26))>>8)
	f.DrawString("存档名字: "+RemoveColorCode(ServerData.LevelName), pt)
	pt = freetype.Pt(20, 275+int(f.PointToFixed(26))>>8)
	f.DrawString("游戏模式: "+GamemodeChinese(ServerData.GameMode), pt)
	pt = freetype.Pt(20, 325+int(f.PointToFixed(26))>>8)
	f.DrawString("连接延迟: "+strconv.FormatInt(ServerData.Delay, 10), pt)

	//将图片写入Buffer
	Buff := bytes.NewBuffer(nil)
	err = png.Encode(Buff, img)
	if err != nil {
		return nil, err
	}
	return Buff, nil
}

/**
 * @description: 移除颜色代码
 * @param {string} String
 * @return {string}
 */
func RemoveColorCode(String string) string {
	reg := regexp.MustCompile(`§[0-9]|§[a-z]`)
	return reg.ReplaceAllString(String, "")
}

/**
 * @description: 汉化游戏模式
 * @param {string} Gamemode
 * @return {string}
 */
func GamemodeChinese(Gamemode string) string {
	switch Gamemode {
	case "Survival":
		return "生存"
	case "Creative":
		return "创造"
	case "Adventure":
		return "冒险"
	default:
		return Gamemode
	}
}
