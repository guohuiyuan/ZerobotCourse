# ZerobotCourse
zerobot框架教学入门

## 为什么选择zerobot作为qqbot开发框架
1. go语言本身的优势
2. zerobot结构清楚

本课程适用于一些喜欢编程, 想要开发自己的qq机器人, 最好是有一门编程语言基础的人~~计算机大一新生~~

- [ZerobotCourse](#zerobotcourse)
  - [为什么选择zerobot作为qqbot开发框架](#为什么选择zerobot作为qqbot开发框架)
  - [第零课](#第零课)
    - [开发环境搭建](#开发环境搭建)
    - [需要以下工具](#需要以下工具)
    - [步骤](#步骤)
    - [参考](#参考)
  - [第一课](#第一课)
    - [了解zerobot框架](#了解zerobot框架)
    - [helloworld](#helloworld)
    - [匹配器](#匹配器)
    - [规则](#规则)
    - [cqhttp API](#cqhttp-api)
    - [消息](#消息)
    - [持续对话](#持续对话)

## 第零课

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
    - 配置vscode的go扩展插件, 参考[go插件配置](第零课/vscodeSettings.md)
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


## 第一课

### 了解zerobot框架
在我看, zerobot框架和那些web框架挺相似的, 不过zerobot是用的websocket, 是双向的, 而web框架是的使用http, 是单向.

zerobot的消息类型相当于路由, rule相当于web中间件, handle都对应的消息的处理, 都定义了context储存上下文的消息

### helloworld
输入hello, 机器回复hello world

```
package hello

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	engine := zero.New()
	engine.OnFullMatch("hello").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("hello world!"))
	})
}

```

注意: 所有engine都是在init的过程中注册的

### 匹配器
[匹配器代码](https://github.com/wdvxdr1123/ZeroBot/blob/main/engine.go)

### 规则
[规则代码](https://github.com/wdvxdr1123/ZeroBot/blob/main/rules.go)

### cqhttp API
[cqhttp API代码](https://github.com/wdvxdr1123/ZeroBot/blob/main/api.go)

### 消息
[消息代码](https://github.com/wdvxdr1123/ZeroBot/blob/main/message/message.go)

### 持续对话
[持续对话代码](https://github.com/wdvxdr1123/ZeroBot/blob/main/event_channel.go)