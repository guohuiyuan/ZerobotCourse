# ZerobotCourse
zerobot框架教学入门

本课程适用于一些喜欢编程, 想要开发自己的qq机器人, 最好是有一门编程语言基础的人~~计算机大一学生~~

- [ZerobotCourse](#zerobotcourse)
  - [第零课](#第零课)
    - [环境搭建](#环境搭建)
    - [需要以下工具](#需要以下工具)
    - [步骤:](#步骤)
    - [参考:](#参考)

## 第零课

### 环境搭建

开发机器人, 开发机器人第一步就是搭建环境

本节课和那些机器人搭建视频教学是一样的, 如果单纯想玩机器人也可以学着搭建一下, 搭建完心里就会有底气, '原来机器人就是这玩意'

工欲利其事必先利其器, 有好工具才能进行更加有效率的开发

### 需要以下工具

- [ ] github([fastgithub](https://github.com/dotnetcore/FastGithub/releases/download/2.1.4/fastgithub_win-x64.zip))
- [ ] [git](https://github.com/git-for-windows/git/releases/download/v2.37.1.windows.1/PortableGit-2.37.1-64-bit.7z.exe)
- [ ] [go](https://golang.google.cn/dl/go1.18.4.windows-amd64.msi)
- [ ] [vscode](https://az764295.vo.msecnd.net/stable/b06ae3b2d2dbfe28bca3134cc6be65935cdfea6a/VSCodeSetup-x64-1.69.1.exe)

因为zerobot是go语言写的, 在github上开源的库, 所以上面四个玩意, 怎么样都没法绕开。

### 步骤: 

- 搭建go语言环境
    - 安装go源代码
    - 安装vscode
    - 配置vscode的go扩展插件, 参考[go插件配置](第零课/vscodeSettings.md)
- 安装[go-cqhttp](https://docs.go-cqhttp.org/) (go-cqhttp相当于qq的第三方客户端, 为我们的程序提供qq的接口, 我们调它的接口, 完成发送qq消息, 下载群文件等操作)
- 下载zerobot框架源码, 并运行 (zerobot框架提供go-cqhttp的接口的封装, 便于开发者编写)

### 参考:
1. [zerobot](https://github.com/wdvxdr1123/ZeroBot)
2. [zerobot-plugin](https://github.com/FloatTech/ZeroBot-Plugin)
3. [派蒙Bot](https://github.com/RicheyJang/PaimengBot)
4. [明日方舟抽卡bot](https://github.com/yuanyan3060/SkadiBot)
5. [zerobot-plugin江林版](https://github.com/Jiang-Red/ZeroBot-Plugin)