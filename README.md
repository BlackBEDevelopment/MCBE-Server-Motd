# Minecraft BedrockEdition Server Motd
我的世界基岩版服务器Motd协议的API封装

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
git clone https://github.com/nyancatda/ShionBot.git --recurse-submodules
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
go get -u github.com/BlackBEDevelopment/MCBE-Server-Motd
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
	data := MotdBEAPI.MotdBE(Host)
	fmt.Println(data)
}
```