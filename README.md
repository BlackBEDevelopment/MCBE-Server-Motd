# Minecraft BedrockEdition Server Motd
我的世界服务器Motd协议封装和HTTP API实践

虽然叫"BedrockEdition Server Motd"，但是也支持Java Server哦 : )

## 🛫 使用
### 官方部署站点
[https://motdbe.blackbe.work/](https://motdbe.blackbe.work/)
### API 文档
[BlackBE Wiki](https://wiki.blackbe.xyz/OpenAPI/tools.html#get-https-motdbe-blackbe-xyz-api-%E8%8E%B7%E5%8F%96%E5%9F%BA%E5%B2%A9%E7%89%88%E6%9C%8D%E5%8A%A1%E5%99%A8%E7%9A%84%E7%8A%B6%E6%80%81)

## 🛠️ 部署
``` shell
#解压资源
unzip MCBE-Server-Motd_{{VERSION_OS_ARCH}}.zip

#赋予运行权限
chmod +x ./MCBE-Server-Motd

#启动
./MCBE-Server-Motd -port 8080
```

## ⚙️ 构建
自行构建前需要拥有 Go >= 1.17，yarn等必要依赖

克隆仓库
``` shell
git clone https://github.com/BlackBEDevelopment/MCBE-Server-Motd.git --recurse-submodules
```

构建静态资源
``` shell
#进入前端目录
cd ./fronend

#安装依赖
yarn install

#开始构建
yarn build
```

编译项目
``` shell
#获取依赖包
go mod tidy

#开始编译
go build .
```

## 🎬 引入项目
### 安装
``` shell
go get -u github.com/BlackBEDevelopment/MCBE-Server-Motd/MotdBEAPI
```

### 例子
``` go
package main

import (
	"fmt"

	"github.com/BlackBEDevelopment/MCBE-Server-Motd/MotdBEAPI"
)

func main() {
	Host := "nyan.xyz:19132"
	data, err := MotdBEAPI.MotdBE(Host)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
```

## 📖许可证
项目采用`Mozilla Public License Version 2.0`协议开源

二次修改源代码需要开源修改后的代码，对源代码修改之处需要提供说明文档
