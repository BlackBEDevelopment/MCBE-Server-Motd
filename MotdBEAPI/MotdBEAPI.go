package MotdBEAPI

import (
	"encoding/hex"
	"net"
	"strings"
	"time"
)

type MotdBEInfo struct {
	Status    string `json:"status"`     //服务器状态
	Host      string `json:"host"`       //服务器Host
	Motd      string `json:"motd"`       //Motd信息
	Agreement string `json:"agreement"`  //协议版本
	Version   string `json:"version"`    //支持的游戏版本
	Online    string `json:"online"`     //在线人数
	Max       string `json:"max"`        //最大在线人数
	LevelName string `json:"level_name"` //存档名字
	GameMode  string `json:"gamemode"`   //游戏模式
	Delay     int64  `json:"delay"`      //连接延迟
}

//使用MotdBE协议发起UDP请求，并获得返回
//Host 服务器地址 nyan.xyz:19132
func MotdBE(Host string) *MotdBEInfo {
	if Host == "" {
		MotdInfo := &MotdBEInfo{
			Status: "offline",
		}
		return MotdInfo
	}

	// 创建连接
	socket, err := net.Dial("udp", Host)
	if err != nil {
		MotdInfo := &MotdBEInfo{
			Status: "offline",
		}
		return MotdInfo
	}
	defer socket.Close()
	// 发送数据
	time1 := time.Now().UnixNano() / 1e6 //记录发送时间
	senddata, _ := hex.DecodeString("0100000000240D12D300FFFF00FEFEFEFEFDFDFDFD12345678")
	_, err = socket.Write(senddata)
	if err != nil {
		MotdInfo := &MotdBEInfo{
			Status: "offline",
		}
		return MotdInfo
	}
	// 接收数据
	UDPdata := make([]byte, 4096)
	socket.SetReadDeadline(time.Now().Add(5 * time.Second)) //设置读取五秒超时
	_, err = socket.Read(UDPdata)
	if err != nil {
		MotdInfo := &MotdBEInfo{
			Status: "offline",
		}
		return MotdInfo
	}
	time2 := time.Now().UnixNano() / 1e6 //记录接收时间
	//解析数据
	if err == nil {
		MotdData := strings.Split(string(UDPdata), ";")
		MotdInfo := &MotdBEInfo{
			Status:    "online",
			Host:      Host,
			Motd:      MotdData[1],
			Agreement: MotdData[2],
			Version:   MotdData[3],
			Online:    MotdData[4],
			Max:       MotdData[5],
			LevelName: MotdData[7],
			GameMode:  MotdData[8],
			Delay:     time2 - time1,
		}
		return MotdInfo
	}

	MotdInfo := &MotdBEInfo{
		Status: "offline",
	}
	return MotdInfo
}
