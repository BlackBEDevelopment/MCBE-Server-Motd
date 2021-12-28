package MotdBEAPI

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)


type MotdJavaInfo struct {
	Status    string `json:"status"`    //服务器状态
	Host      string `json:"host"`      //服务器Host
	Motd      string `json:"motd"`      //Motd信息
	Agreement int    `json:"agreement"` //协议版本
	Version   string `json:"version"`   //支持的游戏版本
	Online    int    `json:"online"`    //在线人数
	Max       int    `json:"max"`       //最大在线人数
	Favicon   string `json:"favicon"`   //服务器图标
	Delay     int64  `json:"delay"`     //连接延迟
}

type MotdJavaJson struct {
	Description struct {
		Text string `json:"text"` //服务器说明文本
	} `json:"description"`
	Players struct {
		Max    int `json:"max"`    //服务器最大在线
		Online int `json:"online"` //服务器当前在线
	} `json:"players"`
	Version struct {
		Name     string `json:"name"`     //可用游戏版本
		Protocol int    `json:"protocol"` //服务器协议版本
	} `json:"version"`

	Favicon string `json:"favicon"` //服务器图标
}

func MotdJava(Host string) (MotdJavaInfo, error) {
	//原代码来自 https://github.com/Cryptkeeper/go-minecraftping
	var MotdInfo MotdJavaInfo

	timeout := time.Second * 5 //设置五秒超时
	deadline := time.Now().Add(timeout)

	time1 := time.Now().UnixNano() / 1e6 //记录发送时间
	conn, err := net.DialTimeout("tcp", Host, timeout)
	if err != nil {
		MotdInfo.Status = "offline"
		return MotdInfo, err
	}
	defer conn.Close()

	if err := conn.SetDeadline(deadline); err != nil {
		MotdInfo.Status = "offline"
		return MotdInfo, err
	}

	// 构造并写入握手包，打开连接，然后写入请求包。
	// More information: https://wiki.vg/Server_List_Ping
	var buf bytes.Buffer
	buf.Write([]byte("\x00"))
	//这个版本号似乎无所谓？
	//575表示1.15.1
	putVarInt(&buf, int32(575))

	//分割服务器地址
	HostAddress := strings.Split(Host, ":")
	Address := HostAddress[0]
	Port, _ := strconv.Atoi(HostAddress[1])

	putVarInt(&buf, int32(len(Address)))
	buf.WriteString(Address)
	binary.Write(&buf, binary.BigEndian, uint16(Port))
	putVarInt(&buf, 1)
	//将缓冲区的长度作为uvarint预加
	var out bytes.Buffer
	putVarInt(&out, int32(buf.Len()))
	out.Write(buf.Bytes())
	handshake := out.Bytes()
	conn.Write(handshake)

	var requestPacket = []byte{1, 0}
	conn.Write(requestPacket)

	reader := bufio.NewReader(conn)
	time2 := time.Now().UnixNano() / 1e6 //记录接收时间

	// 读取并计算传入数据包的长度
	_, err = binary.ReadUvarint(reader)
	if err != nil {
		MotdInfo.Status = "offline"
		return MotdInfo, err
	}

	// 读取数据包，并验证ID是否为0
	packetId, err := binary.ReadUvarint(reader)
	if err != nil {
		MotdInfo.Status = "offline"
		return MotdInfo, err
	}
	if packetId != 0 {
		MotdInfo.Status = "offline"
		return MotdInfo, fmt.Errorf("received invalid packetId (expected 0!) %d", packetId)
	}

	// 读取传入JSON负载的长度（作为uvarint）。将以下字节读入缓冲区，然后将[]byte解组到其结构表示响应中。
	length, err := binary.ReadUvarint(reader)
	if err != nil {
		MotdInfo.Status = "offline"
		return MotdInfo, err
	}
	payload := make([]byte, length)
	if _, err = io.ReadFull(reader, payload); err != nil {
		MotdInfo.Status = "offline"
		return MotdInfo, err
	}

	//解析Json
	var resp MotdJavaJson
	if err = json.Unmarshal(payload, &resp); err != nil {
		MotdInfo.Status = "offline"
		return MotdInfo, err
	}

	//对返回进行二次封装
	MotdInfo.Status = "online"
	MotdInfo.Host = Host
	MotdInfo.Motd = resp.Description.Text
	MotdInfo.Agreement = resp.Version.Protocol
	MotdInfo.Version = resp.Version.Name
	MotdInfo.Online = resp.Players.Online
	MotdInfo.Max = resp.Players.Max
	MotdInfo.Favicon = resp.Favicon
	MotdInfo.Delay = time2 - time1 //计算延迟
	return MotdInfo, nil
}

// 分配一个[]byte的二进制缓冲区。MAXVarintern32并将值作为uvarint32写入。修剪并写入buf。
func putVarInt(buf *bytes.Buffer, value int32) {
	bytes := make([]byte, binary.MaxVarintLen32)
	bytesWritten := binary.PutUvarint(bytes, uint64(value))

	buf.Write(bytes[:bytesWritten])
}
