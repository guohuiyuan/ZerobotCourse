package hello

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type user struct {
	Name string `flag:"n"`
	Age  int    `flag:"a"`
}

func init() {
	engine := zero.New()
	engine.OnFullMatch("完全匹配").Handle(func(ctx *zero.Ctx) {
		matched := ctx.State["matched"].(string)
		ctx.SendChain(message.Text("完全匹配的匹配词: ", matched))
	})
	engine.OnFullMatchGroup([]string{"完全匹配组1", "完全匹配组2"}).Handle(func(ctx *zero.Ctx) {
		matched := ctx.State["matched"].(string)
		ctx.SendChain(message.Text("完全匹配组的匹配词: ", matched))
	})
	engine.OnKeyword("关键词").Handle(func(ctx *zero.Ctx) {
		keyword := ctx.State["keyword"].(string)
		ctx.SendChain(message.Text("关键词匹配的关键词: ", keyword))
	})
	engine.OnCommand("命令").Handle(func(ctx *zero.Ctx) {
		command := ctx.State["command"].(string)
		args := ctx.State["args"].(string)
		ctx.SendChain(message.Text("命令匹配的命令: ", command, "\n命令匹配的参数: ", args))
	})
	engine.OnPrefix("前缀").Handle(func(ctx *zero.Ctx) {
		prefix := ctx.State["prefix"].(string)
		args := ctx.State["args"].(string)
		ctx.SendChain(message.Text("前缀匹配的前缀: ", prefix, "\n前缀匹配的参数: ", args))
	})
	engine.OnRegex(`(.*)正在(.*)`).Handle(func(ctx *zero.Ctx) {
		regexMatched := ctx.State["regex_matched"].([]string)
		ctx.SendChain(message.Text("正则匹配的匹配组: ", regexMatched))
	})
	engine.OnSuffix("后缀").Handle(func(ctx *zero.Ctx) {
		suffix := ctx.State["suffix"].(string)
		args := ctx.State["args"].(string)
		ctx.SendChain(message.Text("后缀匹配的后缀: ", suffix, "\n后缀匹配的参数: ", args))
	})
	engine.OnShell("用户", user{}).Handle(func(ctx *zero.Ctx) {
		args := ctx.State["args"].([]string)
		u := ctx.State["flag"].(*user)
		ctx.SendChain(message.Text("shell匹配的结构体: ", u, "\nshell匹配的参数: ", args))
	})
}
