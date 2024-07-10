// MotdBE协议封装
package MotdBEAPI

import (
	"encoding/binary"
	"net"
	"strconv"
	"strings"
	"time"
)

// MotdBE信息
type MotdBEInfo struct {
	Status         string `json:"status"`           //服务器状态 (online/offline)
	Host           string `json:"host"`             //服务器Host
	Motd           string `json:"motd"`             //Motd信息
	Agreement      int    `json:"agreement"`        //协议版本
	Version        string `json:"version"`          //支持的游戏版本
	Online         int    `json:"online"`           //在线人数
	Max            int    `json:"max"`              //最大在线人数
	LevelName      string `json:"level_name"`       //存档名字
	GameMode       string `json:"gamemode"`         //游戏模式
	ServerUniqueID int    `json:"server_unique_id"` //服务器唯一ID
	Delay          int64  `json:"delay"`            //连接延迟
}

/**
 * @description: 通过UDP请求获取MotdBE信息
 * @param {string} Host 服务器地址，nyan.xyz:19132
 * @return {*MotdBEInfo} MotdBE信息
 * @return {error} 错误信息
 */
func MotdBE(Host string) (*MotdBEInfo, error) {
	errorReturn := &MotdBEInfo{
		Status: "offline",
	}

	if Host == "" {
		return errorReturn, nil
	}

	// 创建连接
	Socket, err := net.Dial("udp", Host)
	if err != nil {
		return errorReturn, err
	}
	defer Socket.Close()

	// 组成发送数据
	PacketID := []byte{0x01} // Packet ID
	// 获取当前时间戳
	ClientSendTime := make([]byte, 8) // 客户端发送时间
	binary.BigEndian.PutUint64(ClientSendTime, uint64(time.Now().Unix()))
	Magic := []byte{0x00, 0xFF, 0xFF, 0x00, 0xFE, 0xFE, 0xFE, 0xFE, 0xFD, 0xFD, 0xFD, 0xFD} // Magic Number
	ClientID := []byte{0x00, 0x00, 0x00, 0x00}                                              // 客户端ID
	// 组合数据
	SendData := append(PacketID, ClientSendTime...)
	SendData = append(SendData, Magic...)
	SendData = append(SendData, ClientID...)

	// 发送数据
	StartTime := time.Now().UnixNano() / 1e6 // 记录发送时间
	_, err = Socket.Write(SendData)
	if err != nil {
		return errorReturn, err
	}

	// 接收数据
	UDPdata := make([]byte, 256)
	Socket.SetReadDeadline(time.Now().Add(5 * time.Second)) // 设置读取五秒超时
	// 读取数据
	_, err = Socket.Read(UDPdata)
	if err != nil {
		return errorReturn, err
	}
	EndTime := time.Now().UnixNano() / 1e6 // 记录接收时间

	// 读取数据
	// PacketID = UDPdata[:1]         // Packet ID
	// ServerSendTime := UDPdata[1:9] // 服务器发送时间
	// ServerGUID := UDPdata[9:17]    // 服务器GUID
	// Magic = UDPdata[17:33]         // Magic Number
	ServerInfo := UDPdata[33:] // 服务器信息

	// 按;分割数据
	MotdData := strings.Split(string(ServerInfo), ";")

	// 解析数据
	MOTD1 := MotdData[1]           // 服务器MOTD line 1
	ProtocolVersion := MotdData[2] // 协议版本
	VersionName := MotdData[3]     // 服务器游戏版本
	PlayerCount := MotdData[4]     // 在线人数
	MaxPlayerCount := MotdData[5]  // 最大在线人数
	ServerUniqueID := MotdData[6]  // 服务器唯一ID
	MOTD2 := MotdData[7]           // 服务器MOTD line 2
	GameMode := MotdData[8]        // 游戏模式
	// GameModeNumeric := MotdData[9]                   // 游戏模式数字

	// 转换数据
	ProtocolVersionInt, err := strconv.Atoi(ProtocolVersion)
	if err != nil {
		return errorReturn, err
	}
	PlayerCountInt, err := strconv.Atoi(PlayerCount)
	if err != nil {
		return errorReturn, err
	}
	MaxPlayerCountInt, err := strconv.Atoi(MaxPlayerCount)
	if err != nil {
		return errorReturn, err
	}
	ServerUniqueIDInt, err := strconv.Atoi(ServerUniqueID)
	if err != nil {
		return errorReturn, err
	}

	MotdInfo := &MotdBEInfo{
		Status:         "online",
		Host:           Host,
		Motd:           MOTD1,
		Agreement:      ProtocolVersionInt,
		Version:        VersionName,
		Online:         PlayerCountInt,
		Max:            MaxPlayerCountInt,
		LevelName:      MOTD2,
		GameMode:       GameMode,
		ServerUniqueID: ServerUniqueIDInt,
		Delay:          EndTime - StartTime,
	}
	return MotdInfo, nil
}
