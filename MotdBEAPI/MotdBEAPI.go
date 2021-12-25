package MotdBEAPI

import (
	"encoding/hex"
	"net"
	"strings"
	"time"

	"github.com/iancoleman/orderedmap"
)

//MotdBE API
//Host 服务器地址 nyan.xyz:19132
func MotdBE(Host string) *orderedmap.OrderedMap {
	Jsondata := orderedmap.New()
	if Host == "" {
		Jsondata.Set("status", "offline")
		return Jsondata
	}

	// 创建连接
	socket, err := net.Dial("udp", Host)
	if err != nil {
		Jsondata.Set("status", "offline")
		return Jsondata
	}
	defer socket.Close()
	// 发送数据
	time1 := time.Now().UnixNano() / 1e6 //记录发送时间
	senddata, _ := hex.DecodeString("0100000000240D12D300FFFF00FEFEFEFEFDFDFDFD12345678")
	_, err = socket.Write(senddata)
	if err != nil {
		Jsondata.Set("status", "offline")
		return Jsondata
	}
	// 接收数据
	UDPdata := make([]byte, 4096)
	socket.SetReadDeadline(time.Now().Add(5 * time.Second)) //设置读取五秒超时
	_, err = socket.Read(UDPdata)
	if err != nil {
		Jsondata.Set("status", "offline")
		return Jsondata
	}
	time2 := time.Now().UnixNano() / 1e6 //记录接收时间
	//解析数据
	if err == nil {
		MotdData := strings.Split(string(UDPdata), ";")
		Jsondata.Set("status", "online")
		Jsondata.Set("host", Host)              //服务器Host
		Jsondata.Set("motd", MotdData[1])       //Motd
		Jsondata.Set("agreement", MotdData[2])  //协议版本
		Jsondata.Set("version", MotdData[3])    //游戏版本
		Jsondata.Set("online", MotdData[4])     //在线人数
		Jsondata.Set("max", MotdData[5])        //最大在线人数
		Jsondata.Set("level_name", MotdData[7]) //存档名字
		Jsondata.Set("gamemode", MotdData[8])   //游戏模式
		Jsondata.Set("delay", time2-time1)      //连接延迟
	}

	return Jsondata
}
