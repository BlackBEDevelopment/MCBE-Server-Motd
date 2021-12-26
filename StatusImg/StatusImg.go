/*
 * @Author: NyanCatda
 * @Date: 2021-12-26 21:23:59
 * @LastEditTime: 2021-12-27 00:59:40
 * @LastEditors: NyanCatda
 * @Description: 服务器状态图片生成
 * @FilePath: \MotdBE\StatusImg\StatusImg.go
 */
package StatusImg

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/BlackBEDevelopment/MCBE-Server-Motd/MotdBEAPI"
	"github.com/golang/freetype"
)

func ServerStatusImg(Host string) *bytes.Buffer {
	//获取服务器信息
	ServerData := MotdBEAPI.MotdBE(Host)
	if ServerData.Status == "offline" {
		offlineImgFile, err := os.Open("StatusImg/background.png")
		if err != nil {
			fmt.Println(err)
		}
		offlineImg, err := png.Decode(offlineImgFile)
		if err != nil {
			fmt.Println(err)
		}
		//将图片写入Buffer
		Buff := bytes.NewBuffer(nil)
		err = png.Encode(Buff, offlineImg)
		if err != nil {
			fmt.Println(err)
		}
		return Buff
	}

	//读取背景图片
	backgroundFile, err := os.Open("StatusImg/background.png")
	if err != nil {
		fmt.Println(err)
	}
	backgroundImg, err := png.Decode(backgroundFile)
	if err != nil {
		fmt.Println(err)
	}

	//转换类型
	img := backgroundImg.(*image.NRGBA)

	//读取字体数据
	fontBytes, err := ioutil.ReadFile("StatusImg/SourceHanSansCN-VF.ttf")
	if err != nil {
		fmt.Println(err)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println(err)
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

	//设置字体的位置
	pt := freetype.Pt(10, 40+int(f.PointToFixed(26))>>8)
	f.DrawString("MOTD: "+RemoveColorCode(ServerData.Motd), pt)
	pt = freetype.Pt(10, 90+int(f.PointToFixed(26))>>8)
	f.DrawString("协议版本: "+ServerData.Agreement, pt)
	pt = freetype.Pt(10, 140+int(f.PointToFixed(26))>>8)
	f.DrawString("游戏版本: "+ServerData.Version, pt)
	pt = freetype.Pt(10, 190+int(f.PointToFixed(26))>>8)
	f.DrawString("在线人数: "+strconv.Itoa(ServerData.Online)+"/"+strconv.Itoa(ServerData.Max), pt)
	pt = freetype.Pt(10, 240+int(f.PointToFixed(26))>>8)
	f.DrawString("存档名字: "+RemoveColorCode(ServerData.LevelName), pt)
	pt = freetype.Pt(10, 290+int(f.PointToFixed(26))>>8)
	f.DrawString("游戏模式: "+GamemodeChinese(ServerData.GameMode), pt)
	pt = freetype.Pt(10, 340+int(f.PointToFixed(26))>>8)
	f.DrawString("连接延迟: "+strconv.FormatInt(ServerData.Delay, 10), pt)

	//将图片写入Buffer
	Buff := bytes.NewBuffer(nil)
	err = png.Encode(Buff, img)
	if err != nil {
		fmt.Println(err)
	}
	return Buff
}

/**
 * @description: 移除颜色代码
 * @param {string} String
 * @return {string}
 */
func RemoveColorCode(String string) string {
	String = strings.Replace(String, "§0", "", -1)
	String = strings.Replace(String, "§1", "", -1)
	String = strings.Replace(String, "§2", "", -1)
	String = strings.Replace(String, "§3", "", -1)
	String = strings.Replace(String, "§4", "", -1)
	String = strings.Replace(String, "§5", "", -1)
	String = strings.Replace(String, "§6", "", -1)
	String = strings.Replace(String, "§7", "", -1)
	String = strings.Replace(String, "§8", "", -1)
	String = strings.Replace(String, "§9", "", -1)
	String = strings.Replace(String, "§a", "", -1)
	String = strings.Replace(String, "§b", "", -1)
	String = strings.Replace(String, "§c", "", -1)
	String = strings.Replace(String, "§d", "", -1)
	String = strings.Replace(String, "§e", "", -1)
	String = strings.Replace(String, "§f", "", -1)
	String = strings.Replace(String, "§g", "", -1)
	return String
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
