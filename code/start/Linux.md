## Linux环境QQBot的部署
+ 框架：NapCat+ZeroBot-Plugin
+ 环境：Ubuntu22.04
+ Shell:  WindTerm

### 介绍
[NapCat](https://napneko.github.io/guide/napcat)是一个Onebot实现

[ZeroBot-Plugin](https://www.github-zh.com/projects/286530881-zerobot-plugin)是一个ZeroBot框架写的插件集

[WindTerm](https://www.github-zh.com/projects/214011414-windterm)是一个SSH连接工具

### 安装
```bash
curl -o napcat.sh https://nclatest.znin.net/NapNeko/NapCat-Installer/main/script/install.sh && sudo bash napcat.sh
```

### 配置文件
NapCat配置的文件目录

![](https://cdn.nlark.com/yuque/0/2025/png/22244395/1743859104337-d7038101-ca76-4e72-941d-3d0634c30266.png)

NapCat的onebot11_114515.json配置文件

```plain
{
  "network": {
    "httpServers": [],
    "httpSseServers": [],
    "httpClients": [],
    "websocketServers": [ {
        "name": "WsServer",// 名字不能重复 唯一标识
        "enable": true,//启用状态
        "host": "0.0.0.0",// 监听主机
        "port": 3001,// 监听端口
        "messagePostFormat": "array",// 消息上报格式 string/array
        "reportSelfMessage": false,// 是否上报自身消息
        "token": "",// 鉴权密钥
        "enableForcePushEvent": true,// 暂时没有作用
        "debug": false,// raw数据上报
        "heartInterval": 30000,// 心跳周期
      }],
    "websocketClients": [],
    "plugins": []
  },
  "musicSignUrl": "",
  "enableLocalFile2Url": false,
  "parseMultMsg": false
}
```

zbp来自Windows环境的编译，或者去Github下载[zbp-amd64-linux](https://github.com/FloatTech/ZeroBot-Plugin/releases/download/v1.9.6/zbp_linux_amd64.tar.gz)的版本，上传到Linux服务器

```plain
$env:GOOS="linux";$env:GOARCH="amd64";go build -ldflags="-s -w" -o zbp
```

zbp文件目录

![](https://cdn.nlark.com/yuque/0/2025/png/22244395/1743858942814-a0a51d7e-dff8-4421-b461-f7eeccf5a050.png)

zbp的config.json配置文件

```plain
{
  "zero": {
    "nickname": ["椛椛", "ATRI", "atri", "亚托莉", "アトリ"],
    "command_prefix": "/",
    "super_users": [114515],
    "ring_len": 4096,
    "latency": 1000000,
    "max_process_time": 240000000000
  },
  "ws": [{ "Url": "ws://127.0.0.1:3001", "AccessToken": "" }]
}
```

### 运行
```bash
screen -dmS napcat bash -c "xvfb-run -a qq --no-sandbox -q QQ号码"
screen -dmS zbp bash -c "/root/zbp -c /root/config.json"
```

+ `screen -dmS napcat bash -c "xvfb-run -a qq --no-sandbox -q QQ号码"`：在前面启动命令基础上，多了 `-q QQ号码` 选项，推测是实现快速登录功能，将指定的 QQ 号码作为参数传入程序。
+ `screen -r napcat`：用于恢复名为 `napcat` 的 `screen` 后台会话，进入到后台运行的程序交互界面。在进入会话后，可通过 `Ctrl + a + d` 再次分离会话，回到原终端，且后台程序不会关闭。 
+ `screen -S napcat -X quit`：这条命令用于关闭名为 `napcat` 的 `screen` 会话，从而停止在该会话中运行的程序。

