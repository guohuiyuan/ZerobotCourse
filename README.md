# ZerobotCourse
ZeroBot框架教学 

qq学习群: 861901070

重新整理了一次代码, 仓库所有代码在code目录下

[qq机器人视频教程](https://www.bilibili.com/video/BV1TF411N71C)

## 为什么选择ZeroBot作为qqbot开发框架
1. go语言本身的优势
   
2. ZeroBot结构清楚

本课程适用于一些喜欢编程, 想要开发自己的qq机器人, 最好是有一门编程语言基础的人~~计算机大一新生~~

- [ZerobotCourse](#zerobotcourse)
	- [为什么选择ZeroBot作为qqbot开发框架](#为什么选择zerobot作为qqbot开发框架)
	- [第0课-golang开发环境搭建](#第0课-golang开发环境搭建)
		- [开发环境搭建](#开发环境搭建)
		- [需要以下工具](#需要以下工具)
		- [步骤](#步骤)
		- [参考](#参考)
	- [番外-linux安装go-cqhttp和zbp](#番外-linux安装go-cqhttp和zbp)
		- [linux一些有用的参考](#linux一些有用的参考)
		- [linux安装命令](#linux安装命令)
		- [linux配置](#linux配置)
	- [第1课-ZeroBot框架以及示例插件](#第1课-zerobot框架以及示例插件)
		- [了解ZeroBot框架](#了解zerobot框架)
		- [helloworld](#helloworld)
		- [示例代码](#示例代码)
		- [匹配器](#匹配器)
		- [规则](#规则)
		- [cqhttp API](#cqhttp-api)
		- [消息](#消息)
		- [持续对话](#持续对话)
	- [第2课-api使用](#第2课-api使用)
		- [http + gjson解析](#http--gjson解析)
		- [http + json.Unmarshal 解析](#http--jsonunmarshal-解析)
		- [gjson](#gjson)
		- [api实践](#api实践)
			- [apifox使用](#apifox使用)
			- [练习](#练习)
			- [示例](#示例)
	- [插播-playwright](#插播-playwright)
	- [第3课-数据库使用](#第3课-数据库使用)
		- [sqlite](#sqlite)
		- [数据库框架](#数据库框架)
		- [FloatTech/sqlite](#floattechsqlite)
		- [gorm](#gorm)
		- [小作文示例代码](#小作文示例代码)
	- [第4课-gg库绘图](#第4课-gg库绘图)
		- [如何学习第三方库](#如何学习第三方库)
		- [10进制数转为rbg格式](#10进制数转为rbg格式)
		- [通过gg库的单元测试, 学习gg库的使用](#通过gg库的单元测试-学习gg库的使用)
		- [gg库示例](#gg库示例)
			- [透明图片](#透明图片)
			- [白色背景](#白色背景)
			- [加载图片](#加载图片)
			- [写字](#写字)
			- [长方形](#长方形)
			- [裁剪图像](#裁剪图像)
			- [倒转](#倒转)
			- [矩阵变换](#矩阵变换)
			- [图像混合颜色](#图像混合颜色)
	- [第5课-制作表情包](#第5课-制作表情包)
		- [步骤](#步骤-1)
		- [例子-终极猎手](#例子-终极猎手)
	- [第6课-爬虫教学](#第6课-爬虫教学)
		- [示例](#示例-1)
		- [xpath, goquery和正则性能测试对比](#xpath-goquery和正则性能测试对比)
		- [使用colly爬取新闻 (简易)](#使用colly爬取新闻-简易)
		- [使用colly爬取日语学习资料 (稍难)](#使用colly爬取日语学习资料-稍难)
		- [小埋写的贴吧图片爬虫](#小埋写的贴吧图片爬虫)
	- [第7课-轮询](#第7课-轮询)
		- [轮询定义](#轮询定义)
		- [场景](#场景)
		- [b站推送](#b站推送)
		- [huggingface](#huggingface)
		- [核心代码](#核心代码)
	- [第8课-api服务器搭建](#第8课-api服务器搭建)
		- [环境准备](#环境准备)
			- [protobuf](#protobuf)
			- [kratos脚手架](#kratos脚手架)
		- [编写proto, 自动生成代码](#编写proto-自动生成代码)
		- [安装swagger](#安装swagger)
		- [音乐下载](#音乐下载)
	- [TodoList](#todolist)

## 第0课-golang开发环境搭建

### 开发环境搭建

开发机器人, 开发机器人第一步就是搭建开发环境

本节课和那些机器人搭建视频教学是一样的, 如果单纯想玩机器人也可以学着搭建一下, 搭建完心里就会有底气, '原来机器人就是这玩意'

工欲利其事必先利其器, 有好工具才能进行更加有效率的开发

### 需要以下工具

- [ ] github([fastgithub](https://github.com/dotnetcore/FastGithub/releases/download/2.1.4/fastgithub_win-x64.zip))
- [ ] [git](https://github.com/git-for-windows/git/releases/download/v2.37.1.windows.1/PortableGit-2.37.1-64-bit.7z.exe)
- [ ] [go](https://golang.google.cn/dl/go1.18.4.windows-amd64.msi)
- [ ] [vscode](https://az764295.vo.msecnd.net/stable/b06ae3b2d2dbfe28bca3134cc6be65935cdfea6a/VSCodeSetup-x64-1.69.1.exe)

因为ZeroBot是go语言写的, 在github上开源的库, 所以上面四个玩意, 怎么样都没法绕开。

### 步骤

- 运行fastgithub
- 搭建go语言环境
    - 安装go源代码
    - 配置go的变量GOROOT和GOPATH, 并将%GOROOT%/bin和%GOPATH%/bin 添加到环境变量
    - go env -w GOPROXY=https://goproxy.cn,direct 修改gomod包镜像
    - 安装vscode
    - 配置vscode的go扩展插件, 参考[go插件配置](第0课/vscodeSettings.md)
- git配置 (用户名和邮箱使用你自己的)
    - 创建github账号
    - ssh-keygen -t rsa -C "1156544355@qq.com"
    - git config --global user.name "小锅饭"
    - git config --global user.email "1156544355@qq.com"
- 安装[go-cqhttp](https://docs.go-cqhttp.org/) (go-cqhttp相当于qq的第三方客户端, 为我们的程序提供qq的接口, 我们调它的接口, 完成发送qq消息, 下载群文件等操作)
- 下载ZeroBot框架源码, 并运行 (ZeroBot框架提供go-cqhttp的接口的封装, 便于开发者编写)

### 参考
1. [ZeroBot](https://github.com/wdvxdr1123/ZeroBot)
2. [ZeroBot-Plugin](https://github.com/FloatTech/ZeroBot-Plugin)
3. [派蒙Bot](https://github.com/RicheyJang/PaimengBot)
4. [明日方舟抽卡bot](https://github.com/yuanyan3060/SkadiBot)
5. [ZeroBot-Plugin江林版](https://github.com/Jiang-Red/ZeroBot-Plugin)

## 番外-linux安装go-cqhttp和zbp
通过文件加速网站, 拿到releases的文件中的下载地址, 然后下载, 解压, 修改权限, 运行

### linux一些有用的参考

[查看linux机器架构](https://www.cnblogs.com/linuxprobe/p/5643668.html)

[GitHub 文件加速](https://gh.api.99988866.xyz/)

[压缩解压命令](https://blog.csdn.net/m0_37824357/article/details/124674008)

[vim乱码](https://blog.csdn.net/qq_17576885/article/details/121975702)

[ubuntu更换镜像源](https://mirrors.tuna.tsinghua.edu.cn/help/ubuntu/)

### linux安装命令

```
# linux ubuntu 
cp -pv /etc/apt/sources.list /etc/apt/sources.list.bak

# 可以搜索清华镜像源, 根据不同版本的ubuntu更换镜像

sed -i -e 's/security.ubuntu.com/mirrors.aliyun.com/g' -e 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list # 更换镜像源

apt-get -y update
apt-get install -y wget unzip gzip curl  # 预下载工具

wget https://gh.api.99988866.xyz/https://github.com/Mrs4s/go-cqhttp/releases/download/v1.0.0-rc3/go-cqhttp_linux_amd64.tar.gz # 下载go-cq

# curl 下载也行
# curl -OL https://gh.api.99988866.xyz/https://github.com/Mrs4s/go-cqhttp/releases/download/v1.0.0-rc3/go-cqhttp_linux_amd64.tar.gz 


mkdir go-cqhttp

tar -xzvf go-cqhttp_linux_amd64.tar.gz -C go-cqhttp

chmod +x go-cqhttp # 添加执行权限

cd go-cqhttp

./go-cqhttp

vim config.yml # 修改配置文件

nohup ./go-cqhttp & # 启动挂后台

tail -n 40 nohup.out # 查看日志

cd .. # 返回上级目录

wget https://gh.api.99988866.xyz/https://github.com/FloatTech/ZeroBot-Plugin/releases/download/v1.5.1-beta4/zbp_linux_amd64.tar.gz # 下载zbp

mkdir zbp

tar -xzvf zbp_linux_amd64.tar.gz -C zbp

cd zbp 

chmod +x zbp # 添加执行权限

./zbp -s config.json # 生成配置文件

vim config.json  #修改配置文件

nohup ./zbp -c config.json & # 带配置文件启动, 启动挂后台

tail -f nohup.out # 看日志 
```

### linux配置
[linux配置](/code/start/Linux.md)

## 第1课-ZeroBot框架以及示例插件

本节课主要讲的是ZeroBot框架的调用和插件的添加

### 了解ZeroBot框架
在我看, ZeroBot框架和那些web框架挺相似的, 不过ZeroBot是用的websocket, 是双向的, 而web框架是的使用http, 是单向.

ZeroBot的消息类型, rule相当于路由, zb的前置和后置处理相当于web中间件, handle都对应的消息的处理, 都定义了context储存上下文的消息

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

func main() {
	zero.RunAndBlock(&zero.Config{
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
进入[第一课](./code/first/main.go)目录, 然后go run main.go

```
cd code/first
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


## 第2课-api使用

本节课主要讲的是api的使用

### http + gjson解析
[gjson](code/second/gjson/main.go)

```
package main

import (
	"fmt"

	"github.com/FloatTech/floatbox/web"
	"github.com/tidwall/gjson"
)

var (
	requestURL = "https://api.shadiao.app/chp"
)

func main() {
	data, err := web.RequestDataWith(web.NewDefaultClient(), requestURL, "GET", "", web.RandUA(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	fmt.Println(gjson.GetBytes(data, "data.text").String())
}

```

### http + json.Unmarshal 解析
[json](code/second/json/main.go)

```
package main

import (
	"encoding/json"
	"fmt"

	"github.com/FloatTech/floatbox/web"
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
	data, err := web.RequestDataWith(web.NewDefaultClient(), requestURL, "GET", "", web.RandUA(), nil)
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

练习素材

[有趣的网站](./练习素材/有趣的网站.md)

[B站2查数据获取源码](./练习素材/B站2查数据获取源码.txt)

#### 示例
[tiangou](code/second/tiangou/main.go)

```
// Package main 舔狗日记
package main

import (
	"fmt"

	"github.com/FloatTech/floatbox/web"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	tiangouURL = "http://bh.ayud.top/tg?qq=%v"
)

func init() {
	engine := zero.New()
	engine.OnRegex(`^舔狗\s?(\d+)$`).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		regexMatched := ctx.State["regex_matched"].([]string)
		ctx.SendChain(message.Text(regexMatched[1]))
		data, err := web.GetData(fmt.Sprintf(tiangouURL, regexMatched[1]))
		if err != nil {
			ctx.SendChain(message.Text("ERROR:", err))
			return
		}
		ctx.SendChain(message.ImageBytes(data))
	})
}

func main() {
	zero.RunAndBlock(&zero.Config{
		NickName:      []string{"bot"},
		CommandPrefix: "/",
		SuperUsers:    []int64{123456},
		Driver: []zero.Driver{
			driver.NewWebSocketClient("ws://127.0.0.1:6700/", ""),
		},
	}, nil)
}


```

## 插播-playwright

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


## 第3课-数据库使用

本节课主要数据库的使用

### sqlite
zbp使用的数据库是sqlite

sqlite是一个进程内的库，实现了自给自足的、无服务器的、零配置的、事务性的 SQL 数据库引擎。它是一个零配置的数据库，这意味着与其他数据库不一样，您不需要在系统中配置。

简单就是本地数据库,sqlite让zbp保持数据库的零依赖

### 数据库框架

zbp目前存在两个数据库框架,一种是自用简易框架---FloatTech/sqlite,一种是gorm

我墙裂推荐gorm

### FloatTech/sqlite

[示例代码](https://github.com/FloatTech/ZeroBot-Plugin/blob/master/plugin/book_review/model.go)

### gorm

gorm需要自己封装一个类型

[示例代码](https://github.com/FloatTech/ZeroBot-Plugin/blob/master/plugin/sleep_manage/model.go)

[学习教程](https://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/gorm/gorm%E7%94%A8%E6%B3%95%E4%BB%8B%E7%BB%8D.html)

### 小作文示例代码
[article](code/third/main.go)

```
package main

import (
	"fmt"
	"time"

	sql "github.com/FloatTech/sqlite"
)

type article struct {
	ID         int64  `db:"id"`
	Title      string `db:"title"`
	Author     string `db:"author"`
	CreateTime string `db:"createTime"`
	Content    string `db:"content"`
}

var db = &sql.Sqlite{}

// 暂时随机选择一个小作文
func getArticleByKeyword(keyword string) (a article) {
	_ = db.Find("main", &a, "where content LIKE '%"+keyword+"%'")
	return
}

func getRandomArticle() (a article) {
	_ = db.Pick("main", &a)
	return
}

func main() {
	db.DBPath = "小作文.db"
	err := db.Open(time.Hour * 24)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Create("main", &article{})
	if err != nil {
		fmt.Println(err)
	}
	n, err := db.Count("main")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("一共有", n, "条记录")
	a := getRandomArticle()
	fmt.Printf("%+v\n", a)
}

```

## 第4课-gg库绘图

### 如何学习第三方库
因为zbp不接入浏览器渲染, 所以只能使用2d进行画图, 而我们最常使用的2d画图库就是gg库

所以本节课讲gg库绘图

使用gg库画图既是技术活更是体力活, 需要你定位所有图像的位置, 非常折磨

一般我们学习go第三方库, 最快的方法就是看第三方库的作者写的example, 那可能有人要问了为啥不去看别人写的技术博客?
有两个原因, 第一, 找技术博客比较麻烦, 第二, 有些库冷到根本没有技术文档, 需要我们多多阅读源码

补充一句, 看源码的方式学习, 不适合初学者, 因为比较难看懂, 没错我说的就是rust, 所以最好是看视频, 还有文档。

### 10进制数转为rbg格式

做bilibili查成分的时候碰到了一个问题就是把10进制数转为rbg格式
[color](code/fourth/color/main.go)

```
package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

var (
	testcolor = 1725515
)

func main() {
	fmt.Print("int2rbg测试结果:")
	fmt.Println(int2rbg(int64(testcolor)))
	colorHex := strconv.FormatInt(int64(testcolor), 16)
	fmt.Println("十六进制测试结果:", colorHex)
	fmt.Print("parseHexColor测试结果:")
	fmt.Println(parseHexColor(colorHex))
}

func int2rbg(t int64) (int64, int64, int64) {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], uint64(t))
	b, g, r := int64(buf[0]), int64(buf[1]), int64(buf[2])
	return r, g, b
}

func parseHexColor(x string) (r, g, b, a int) {
	x = strings.TrimPrefix(x, "#")
	a = 255
	if len(x) == 3 {
		format := "%1x%1x%1x"
		fmt.Sscanf(x, format, &r, &g, &b)
		r |= r << 4
		g |= g << 4
		b |= b << 4
	}
	if len(x) == 6 {
		format := "%02x%02x%02x"
		fmt.Sscanf(x, format, &r, &g, &b)
	}
	if len(x) == 8 {
		format := "%02x%02x%02x%02x"
		fmt.Sscanf(x, format, &r, &g, &b, &a)
	}
	return
}
```

```
func (dc *Context) SetHexColor(x string) {
	r, g, b, a := parseHexColor(x)
	dc.SetRGBA255(r, g, b, a)
}
```

自己做了一个int2rbg, 后面了解gg库里处理16进制数, 有一个parseHexColor方法, 测试证明这两种转化方式一样

### 通过gg库的单元测试, 学习gg库的使用
gg库的单元测试是在[context_test.go](https://github.com/fogleman/gg/blob/master/context_test.go)中的

### gg库示例

推荐阅读

[go原生image库画图](https://www.cnblogs.com/Finley/p/16589798.html)

[go字符画](https://blog.csdn.net/ccboy2009/article/details/77803471)

从一个地方看到的, 感觉指针可以直接转

将image.Image转换成 *image.RGBA
```
func ImageToRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
	draw.Draw(dst, bounds, src, bounds.Min, draw.Src)
	return dst
}
```

#### 透明图片
[blank](code/fourth/blank/main.go)

```
package main

import "github.com/fogleman/gg"

func main() {
	dc := gg.NewContext(100, 100)
	dc.SavePNG("out.png")
}
```

#### 白色背景
[white](code/fourth/white/main.go)

```
package main

import (
	"image/color"

	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(200, 100)
	dc.SetColor(color.White)
	dc.Clear()
	dc.SavePNG("out.png")
}

```

#### 加载图片
[load](code/fourth/load/main.go)

```
package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(4000, 2000)
	back, err := gg.LoadImage("原神.jpg")
	if err != nil {
		fmt.Println(err)
	}
	dc.DrawImage(back, 500, 0)
	dc.SavePNG("out.png")
}

```

#### 写字
[drawstring](code/fourth/drawstring/main.go)

```
package main

import (
	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(1024, 1024)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("font/regular-bold.ttf", 50); err != nil {
		panic(err)
	}
	dc.DrawString("你好", 0, 50)
	dc.SavePNG("out.png")
}

```

#### 长方形
[rectangle](code/fourth/rectangle/main.go)

```
package main

import (
	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(1024, 1024)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.DrawRectangle(200, 100, 400, 200)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG("out.png")
}

```

#### 裁剪图像
[cut](code/fourth/cut/main.go)

```
package main

import (
	"fmt"
	"image"
	"os"

	"github.com/FloatTech/imgfactory"
	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(4000, 2000)
	back, err := gg.LoadImage("fourth/原神.jpg")
	if err != nil {
		fmt.Println(err)
	}
	dc.DrawImage(back, 0, 0)
	im := dc.Image().(*image.RGBA)
	nim := im.SubImage(image.Rect(0, 0, 1000, 100))
	f, err := os.Create("out.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = imgfactory.WriteTo(nim, f)
	_ = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}


```

#### 倒转
[invert](code/fourth/invert/main.go)

```
package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(4000, 2000)
	back, err := gg.LoadImage("fourth/原神.jpg")
	if err != nil {
		fmt.Println(err)
	}
	dc.InvertY()
	dc.DrawImage(back, 2000, 0)
	dc.SavePNG("out.png")
}
```

#### 矩阵变换
[transform](code/fourth/transform/main.go)

gg.Scale(0.5,0.5) 意思是让后面的矩阵长宽缩短一半

gg库绕某个点旋转

shearAbout可以把图片变成平行四边形, 用于透视变换

```
package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(4000, 2000)
	back, err := gg.LoadImage("fourth/原神.jpg")
	if err != nil {
		fmt.Println(err)
	}
	dc.Scale(0.5, 0.5)
	dc.ShearAbout(0.5, 0, 2000, 1000)
	// dc.ShearAbout(0, 0.5, 2000, 1000)
	for i := 0; i <= 9; i++ {
		dc.DrawImage(back, 2000, 1000)
		dc.RotateAbout(gg.Radians(float64(10)), 2000, 1000)
	}
	dc.SavePNG("out.png")
}

```

#### 图像混合颜色

```
待定
```

## 第5课-制作表情包

实际就是用代码去重现图片编辑软件的操作

### 步骤

1. 准备背景图片

- [抠图](https://ps.gaoding.com/#/) 

- [~~借鉴现有素材~~](https://github.com/noneplugin/nonebot-plugin-petpet)

- [智能抠图](https://www.remove.bg/zh) 

2. 使用图像库进行图形变换

- 主要使用"github.com/FloatTech/zbputils/img"库进行图像变换
  
[zbp示例](https://github.com/FloatTech/ZeroBot-Plugin/blob/master/plugin/gif/png.go)

```
// roll 滚
func roll(cc *context, value ...string) (string, error) {
	// 下载背景图片到本地
	_ = value
	var wg sync.WaitGroup
	var err error
	var m sync.Mutex
	piclen := 8
	name := cc.usrdir + "roll.gif"
	c := dlrange("roll", piclen, &wg, func(e error) {
		m.Lock()
		err = e
		m.Unlock()
	})
	wg.Wait()
	if err != nil {
		return "", err
	}

	// 加载用户头像
	im, err := img.LoadFirstFrame(cc.headimgsdir[0], 210, 210)
	if err != nil {
		return "", err
	}

	// 定义用户头像的变换
	locs := [][]int{{87, 77, 0}, {96, 85, -45}, {92, 79, -90}, {92, 78, -135}, {92, 75, -180}, {92, 75, -225}, {93, 76, -270}, {90, 80, -315}}

	// 加载背景图片
	imgs, err := loadFirstFrames(c, piclen)
	if err != nil {
		return "", err
	}
	roll := make([]*image.NRGBA, piclen)
	for i := 0; i < piclen; i++ {
		// 实际变换, 在背景图底下插入变换位置的头像,
		roll[i] = imgs[i].InsertBottomC(img.Rotate(im.Im, float64(locs[i][2]), 0, 0).Im, 0, 0, locs[i][0]+105, locs[i][1]+105).Im
	}

	// 返回制作好的gif
	return "file:///" + name, writer.SaveGIF2Path(name, img.MergeGif(7, roll))
}
```

本节课使用gg等图像库制作表情包

### 例子-终极猎手

[素材](第5课/0.png)

[图片定位](https://github.com/Mannix1994/PointsPicker)

```
// zhongjilieshou 终极猎手
func zhongjilieshou(cc *context, args ...string) (string, error) {
	_ = args
	var wg sync.WaitGroup
	var m sync.Mutex
	var err error
	c := dlrange("zhongjilieshou", 1, &wg, func(e error) {
		m.Lock()
		err = e
		m.Unlock()
	})
	wg.Wait()
	if err != nil {
		return "", err
	}
	imgs, err := loadFirstFrames(c, 1)
	if err != nil {
		return "", err
	}
	name := cc.usrdir + "Zhongjilieshou.png"
	im, err := img.LoadFirstFrame(cc.headimgsdir[0], 360, 360)
	if err != nil {
		return "", err
	}
	imgnrgba := imgs[0].InsertBottom(im.Im, 0, 0, 686, 136).Im
	return "file:///" + name, writer.SavePNG2Path(name, imgnrgba)
}
```

## 第6课-爬虫教学

爬虫教学(http, xpath, goquery, 正则, goroutine, 数据库)

主要使用colly框架来写爬虫, 因为如果是我们自己用http写爬虫, 只适合查找站内某些的信息, 不合适爬全站的信息

温馨提示, 千万不要直播搞爬虫, 并打开色情网站, 对, 说的就是asmr

### 示例
[asmr](https://github.com/DiheChen/go-asmr-spider)

[漫画猫](https://github.com/guohuiyuan/maofly-spider)

[爬虫大佬](https://github.com/xianyucoder)

[Golang goquery selector(选择器) 教程](https://blog.csdn.net/songhao8080/article/details/103669861)

### xpath, goquery和正则性能测试对比

[性能对比代码](第6课/compare)

![性能对比结果](https://user-images.githubusercontent.com/54976075/189469802-d006be96-fdb0-40cf-88de-e7147f8ee36d.jpg)

事实证明, 算加载的话, 正则确实比xpath和goquery效率高, 但goquery比正则好用, 而且dom可重复利用啊

```
cd 第6课/compare && go test -bench . -benchmem
```

### 使用colly爬取新闻 (简易)
colly是go语言知名的爬虫框架, 我也是近期学习的, 方便易用, 比较适合爬一些没有加密的网站或者只需要带cookie的网站

也可以爬那种带加密的, 可以自定义加签方法。

colly可以通过colly库里面的example学习

下面是一个爬新闻文章的一个demo

[爬新闻demo](第6课/news)

```
cd 第6课/news && go run .
```
 
### 使用colly爬取日语学习资料 (稍难)

[http://jp.tingroom.com/ 的听力资源](第6课/jpTingroomSpider)

```
cd 第6课/jpTingroomSpider && go run .
```

### 小埋写的贴吧图片爬虫
未使用colly框架, 自己起协程

```
cd 第6课/tieba && go run .
```

## 第7课-轮询

### 轮询定义

轮询(Polling)是指不管服务器端有没有更新,客户端(通常是指浏览器)都定时的发送请求进行查询,轮询的结果可能是服务器端有新的更新过来,也可能什么也没有,只是返回个空的信息

从api的角度来说, 就是定时请求接口, 监视数据变化, 展示这些变换的数据

### 场景

具体场景就是b站推送的实现和huggingface网站api的调用

### b站推送

我之前做了两种类型的b站推送, 一种是动态推送, 一种是直播推送. 

动态推送实现是, 每隔10秒拿到每个up主最近一条动态的发布时间并存入内存, 如果这个时间大于之前储存的更新时间, 就把这条动态推送到qq群.

直播推送实现的, 每隔10秒拿到每个直播间的状态, 如果直播的状态有更新, 且直播状态是由闲置到直播中, 就把直播信息发到qq群

[b站推送代码](https://github.com/FloatTech/ZeroBot-Plugin/blob/master/plugin/bilibili/bilibilipush.go)

### huggingface

huggingface号称ai界的github, 虽然只是demo, 但里面的模型确实好玩.

里面还有api供我们简单体验这些模型的效果, 现在由我教大家如何接入里面的api.

[猫雷vits接入代码](https://github.com/FloatTech/ZeroBot-Plugin/blob/master/plugin/vitsnyaru/vitsnyaru.go)

主要有两个api, 一个push, 一个是status, push就是上传你的文本,图片或者音频等模型的输入, status就是拿到模型的输出.

push只需要请求一次, status需要请求很多次, status是一个异步的请求, 返回里面有一个status字段判断输出是否完成, 很有趣的事, 一旦你拿到了数据, 输出结束, 你再取拿同样的参数请求status, status会报500的错误.

huggingface有些api比较特殊, 不是https请求, 而是wss请求, wss是双向的请求, 所以不需要轮询, 只需要保持连接, 获得输出发送就行, 参考[MagicPrompt-Stable-Diffusion](https://github.com/FloatTech/ZeroBot-Plugin/blob/master/plugin/magicprompt/magicprompt.go)

### 核心代码

通过轮询, 拿到status的状态, 如果是已完成, 跳出循环, 出现err或超时, 也跳出循环

```
			statusReq = hf.StatusRequest{
				Hash: pushRes.Hash,
			}

			t := time.NewTicker(time.Second * 1)
			defer t.Stop()
		LOOP:
			for {
				select {
				case <-t.C:
					data, err = hf.Status(statusURL, statusReq)
					if err != nil {
						ch <- data
						break LOOP
					}
					if gjson.ParseBytes(data).Get("status").String() == hf.CompleteStatus {
						ch <- data
						break LOOP
					}
				case <-_ctx.Done():
					ch <- data
					break LOOP
				}
			}
```

## 第8课-api服务器搭建

api服务器搭建(kratos,数据库)

~~先搁置, 等什么时候go-cqhttp被毁灭了再写, 到时候大家就只能玩api了~~

最近想写个音乐下载后端, 所以开始写了, 重新学习一下kratos的写法 

[kratos使用](https://www.cnblogs.com/jiujuan/p/16331967.html)

### 环境准备

#### protobuf
首先了解一下什么是protobuf

Google Protocol Buffer( 简称 Protobuf) 是 Google 公司内部的混合语言数据标准。他们用于 RPC 系统和持续数据存储系统。提供一个具有高效的协议数据交换格式工具库(类似Json)。
但相比于Json，Protobuf有更高的转化效率，时间效率和空间效率都是JSON的3-5倍。
可用于通讯协议、数据存储等领域的语言无关、平台无关、可扩展的序列化结构数据格式。目前提供了 C++、Java、Python 、OC、Swift等语言的 API。总只一句话就是很好，支持多平台且与语言无关。

我们想编写protobuf, 通过protobuf-gen-go生成go代码, 编写函数名的烦恼, 我们只要编写函数的实现就行了

下载protobuf-gen-go
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

#### kratos脚手架
下载kratos

```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

kratos新建项目

```
kratos new music-downloader
```

kratos运行

```
kratos run
curl http://localhost:8000/helloworld/kratos
```

### 编写proto, 自动生成代码
添加proto

```
kratos proto add api/helloworld/v1/music.proto
```

```
syntax = "proto3";

package api.helloworld.v1;

import "google/api/annotations.proto";

option go_package = "music-downloader/api/helloworld/v1;v1";
option java_multiple_files = true;
option java_package = "api.helloworld.v1";

service Music {
	rpc SearchMusic (SearchMusicRequest) returns (SearchMusicReply){
		option (google.api.http) = {
			get: "/music/search"
		};
	};
}

message SearchMusicRequest {
	string name = 1;
}

message SearchMusicReply {
	repeated SearchMusicReplyItem musicList = 1;			
}

message SearchMusicReplyItem {
	string url = 1;
	string audio = 2;
	string title = 3;
	string content = 4;
	string image = 5;
}
```

生成基础参数

```
kratos proto client api/helloworld/v1/music.proto
```

添加service层

```
kratos proto server api/helloworld/v1/music.proto -t internal/service
```

修改 internal\server 代码

**internal\server\grpc.go**

```
package server

import (
	v1 "music-downloader/api/helloworld/v1"
	"music-downloader/internal/conf"
	"music-downloader/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, music *service.MusicService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterMusicServer(srv, music)
	return srv
}

```

**internal\server\http.go**

```
package server

import (
	v1 "music-downloader/api/helloworld/v1"
	"music-downloader/internal/conf"
	"music-downloader/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, music *service.MusicService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterMusicHTTPServer(srv, music)
	return srv
}

```

### 安装swagger
生成swagger文档

proto导入太痛苦, 最后我还是把文件复制到third_party里去了

```
git clone git@github.com:grpc-ecosystem/grpc-gateway.git
git clone git@github.com:gogo/protobuf.git
```

突然发现, go应该不适合写swagger, swagger是写在proto里的, 这太麻烦了

### 音乐下载


## TodoList
- [x] qq自动发说说, 参考[OPQBot](https://github.com/opq-osc/OPQBot)
- [ ] 每日番剧, 搜索动漫资源
- [x] huggingface各种有趣模型接入
- [ ] 幻影坦克
- [ ] zbp爬虫优化(使用goquery,colly重构)
- [ ] 微博推送, a站推送
- [ ] musescore音乐下载


