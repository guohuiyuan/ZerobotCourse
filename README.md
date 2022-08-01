# ZerobotCourse
zerobot框架教学入门

## 为什么选择zerobot作为qqbot开发框架
1. go语言本身的优势
2. zerobot结构清楚

本课程适用于一些喜欢编程, 想要开发自己的qq机器人, 最好是有一门编程语言基础的人~~计算机大一新生~~

- [ZerobotCourse](#zerobotcourse)
	- [为什么选择zerobot作为qqbot开发框架](#为什么选择zerobot作为qqbot开发框架)
	- [第0课](#第0课)
		- [开发环境搭建](#开发环境搭建)
		- [需要以下工具](#需要以下工具)
		- [步骤](#步骤)
		- [参考](#参考)
	- [第1课](#第1课)
		- [了解zerobot框架](#了解zerobot框架)
		- [helloworld](#helloworld)
		- [示例代码](#示例代码)
		- [匹配器](#匹配器)
		- [规则](#规则)
		- [cqhttp API](#cqhttp-api)
		- [消息](#消息)
		- [持续对话](#持续对话)
	- [第2课](#第2课)
		- [http + gjson解析](#http--gjson解析)
		- [http + json.Unmarshal 解析](#http--jsonunmarshal-解析)
		- [gjson](#gjson)
		- [api实践](#api实践)
			- [apifox使用](#apifox使用)
			- [练习](#练习)
			- [示例](#示例)
	- [插播](#插播)
	- [第三课](#第三课)
		- [FloatTech/sqlite](#floattechsqlite)
		- [gorm](#gorm)

## 第0课

### 开发环境搭建

开发机器人, 开发机器人第一步就是搭建开发环境

本节课和那些机器人搭建视频教学是一样的, 如果单纯想玩机器人也可以学着搭建一下, 搭建完心里就会有底气, '原来机器人就是这玩意'

工欲利其事必先利其器, 有好工具才能进行更加有效率的开发

### 需要以下工具

- [ ] github([fastgithub](https://github.com/dotnetcore/FastGithub/releases/download/2.1.4/fastgithub_win-x64.zip))
- [ ] [git](https://github.com/git-for-windows/git/releases/download/v2.37.1.windows.1/PortableGit-2.37.1-64-bit.7z.exe)
- [ ] [go](https://golang.google.cn/dl/go1.18.4.windows-amd64.msi)
- [ ] [vscode](https://az764295.vo.msecnd.net/stable/b06ae3b2d2dbfe28bca3134cc6be65935cdfea6a/VSCodeSetup-x64-1.69.1.exe)

因为zerobot是go语言写的, 在github上开源的库, 所以上面四个玩意, 怎么样都没法绕开。

### 步骤

- 运行fastgithub
- 搭建go语言环境
    - 安装go源代码
    - 配置go的变量GOROOT和GOPATH, 并将%GOROOT%/bin和%GOPATH%/bin 添加到环境变量
    - 安装vscode
    - 配置vscode的go扩展插件, 参考[go插件配置](第0课/vscodeSettings.md)
- git配置 (用户名和邮箱使用你自己的)
    - 创建github账号
    - ssh-keygen -t rsa -C "1156544355@qq.com"
    - git config --global user.name "小锅饭"
    - git config --global user.email "1156544355@qq.com"
- 安装[go-cqhttp](https://docs.go-cqhttp.org/) (go-cqhttp相当于qq的第三方客户端, 为我们的程序提供qq的接口, 我们调它的接口, 完成发送qq消息, 下载群文件等操作)
- 下载zerobot框架源码, 并运行 (zerobot框架提供go-cqhttp的接口的封装, 便于开发者编写)

### 参考
1. [zerobot](https://github.com/wdvxdr1123/ZeroBot)
2. [zerobot-plugin](https://github.com/FloatTech/ZeroBot-Plugin)
3. [派蒙Bot](https://github.com/RicheyJang/PaimengBot)
4. [明日方舟抽卡bot](https://github.com/yuanyan3060/SkadiBot)
5. [zerobot-plugin江林版](https://github.com/Jiang-Red/ZeroBot-Plugin)


## 第1课

本节课主要讲的是zerobot框架和第一个插件的添加

### 了解zerobot框架
在我看, zerobot框架和那些web框架挺相似的, 不过zerobot是用的websocket, 是双向的, 而web框架是的使用http, 是单向.

zerobot的消息类型,rule相当于路由, zb的前置和后置处理相当于web中间件, handle都对应的消息的处理, 都定义了context储存上下文的消息

### helloworld
输入hello, 机器回复hello world

```
package main

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	engine := zero.New()
	engine.OnFullMatch("hello").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("hello world!"))
	})
}

func main(){
	zero.RunAndBlock(zero.Config{
		NickName:      []string{"bot"},
		CommandPrefix: "/",
		SuperUsers:    []int64{123456},
		Driver: []zero.Driver{
			driver.NewWebSocketClient("ws://127.0.0.1:6700/", ""),
		},
	}, nil)
}

```

注意: 所有engine都是在init的过程中注册的

### 示例代码
进入[第一课](./第1课/main.go)目录,然后go run main.go

```
cd 第1课
go run main.go
```

### 匹配器
[匹配器代码](https://github.com/wdvxdr1123/ZeroBot/blob/main/engine.go)

### 规则
[规则代码](https://github.com/wdvxdr1123/ZeroBot/blob/main/rules.go)

### cqhttp API
[cqhttp API代码](https://github.com/wdvxdr1123/ZeroBot/blob/main/api.go)

[合并转发参考代码](https://github.com/FloatTech/ZeroBot-Plugin/blob/master/plugin/coser/coser.go)

[群管理相关api参考代码](https://github.com/FloatTech/ZeroBot-Plugin/blob/master/plugin/manager/manager.go)

[ctxext api的包装](https://github.com/FloatTech/zbputils/blob/4e1d708dffe9/ctxext/message.go)

### 消息
[消息代码](https://github.com/wdvxdr1123/ZeroBot/blob/main/message/message.go)

### 持续对话
[持续对话代码](https://github.com/wdvxdr1123/ZeroBot/blob/main/event_channel.go)

[对话参考代码](https://github.com/FloatTech/ZeroBot-Plugin/blob/master/plugin/wordle/wordle.go)


## 第2课

本节课主要讲的是api的使用

### http + gjson解析

```
package main

import (
	"fmt"

	"github.com/FloatTech/zbputils/web"
	"github.com/tidwall/gjson"
)

var (
	requestURL = "https://api.shadiao.app/chp"
)

func main() {
	data, err := web.RequestDataWith(web.NewDefaultClient(), requestURL, "GET", "", web.RandUA())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	fmt.Println(gjson.GetBytes(data, "data.text").String())
}

```

### http + json.Unmarshal 解析

```
package main

import (
	"encoding/json"
	"fmt"

	"github.com/FloatTech/zbputils/web"
)

var (
	requestURL = "https://api.shadiao.app/chp"
)

type shadiao struct {
	Data struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"data"`
}

func main() {
	data, err := web.RequestDataWith(web.NewDefaultClient(), requestURL, "GET", "", web.RandUA())
	if err != nil {
		fmt.Println(err)
		return
	}
	var s shadiao
	json.Unmarshal(data, &s)
	fmt.Printf("%+v\n", s)
}
```

### gjson

[gjson使用教程](https://www.topgoer.com/%E5%85%B6%E4%BB%96/gjson.html)

### api实践

#### apifox使用

使用apifox来测试api是否可用

#### 练习

使用现成的api来编写插件

[ddbot模板使用的api](https://docs.qq.com/doc/DVERZT0FrYVNEb01j)
[图片api整合](./第2课/图片api整合.txt)

#### 示例
+
```
// Package tiangou2 舔狗日记
package tiangou2

import (
	"fmt"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/web"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	tiangouURL = "http://bh.ayud.top/tg?qq=%v"
)

func init() {
	engine := control.Register("tiangou2", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help:             "舔狗日记\n- 舔狗123456",
	})
	engine.OnRegex(`^舔狗\s?(\d+)$`).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		regexMatched := ctx.State["regex_matched"].([]string)
		ctx.SendChain(message.Text(regexMatched[1]))
		data, err := web.GetData(fmt.Sprintf(tiangouURL, regexMatched[1]))
		if err != nil {
			ctx.SendChain(message.Text("Error:", err))
		}
		ctx.SendChain(message.ImageBytes(data))
	})
}

```

## 插播

之前用chromedp截图,发现linux安装chrome真的太消耗资源的,所以弃用了浏览器截图

现在听别人说playwright挺好用,然后我又发现了playwright-go,所以我想用这个库做浏览器截图

go下载playwright

```
go install github.com/playwright-community/playwright-go/cmd/playwright@latest
# centos7适用,centos8用不了
playwright install --with-deps
```

linux安装的时候,如果是centos系统还要安装apt-get

```
curl https://raw.githubusercontent.com/dvershinin/apt-get-centos/master/apt-get.sh -o /usr/local/bin/apt-get

chmod 0755 /usr/local/bin/apt-get
```

[插件代码](https://github.com/FloatTech/ZeroBot-Plugin-Playground/blob/main/playwright/playwright.go)


## 第三课

本节课主要数据库的使用

zbp目前存在两个数据库框架,一种是自用简易框架---FloatTech/sqlite,一种是gorm

我墙裂推荐gorm

### FloatTech/sqlite

[示例代码](https://github.com/FloatTech/ZeroBot-Plugin/blob/master/plugin/book_review/model.go)

### gorm

gorm需要自己封装一个类型

[示例代码](https://github.com/FloatTech/ZeroBot-Plugin/blob/master/plugin/sleep_manage/model.go)

[学习教程](https://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/gorm/gorm%E7%94%A8%E6%B3%95%E4%BB%8B%E7%BB%8D.html)