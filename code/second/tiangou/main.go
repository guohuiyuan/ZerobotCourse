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
