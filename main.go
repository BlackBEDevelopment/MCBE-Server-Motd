package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/iancoleman/orderedmap"
)

func main() {
	Jsondata := orderedmap.New()
	// 创建连接
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(47, 100, 252, 210),
		Port: 19132,
	})
	if err != nil {
		Jsondata.Set("status", "offline")
	}
	defer socket.Close()
	// 发送数据
	time1 := time.Now().UnixNano() / 1e6 //记录发送时间
	senddata, _ := hex.DecodeString("0100000000240D12D300FFFF00FEFEFEFEFDFDFDFD12345678")
	_, err = socket.Write(senddata)
	if err != nil {
		Jsondata.Set("status", "offline")
	}
	// 接收数据
	UDPdata := make([]byte, 4096)
	_, _, err = socket.ReadFromUDP(UDPdata)
	if err != nil {
		Jsondata.Set("status", "offline")
	}
	time2 := time.Now().UnixNano() / 1e6 //记录接收时间
	//解析数据
	MotdData := strings.Split(string(UDPdata), ";")
	if err == nil {
		Jsondata.Set("status", "offline")
		Jsondata.Set("motd", MotdData[1])      //Motd
		Jsondata.Set("agreement", MotdData[2]) //协议版本
		Jsondata.Set("version", MotdData[3])   //游戏版本
		Jsondata.Set("online", MotdData[4])    //在线人数
		Jsondata.Set("max", MotdData[5])       //最大在线人数
		Jsondata.Set("gamemode", MotdData[6])  //游戏模式
		Jsondata.Set("delay", time2-time1)     //连接延迟
	}

	fmt.Println(Jsondata)
}
