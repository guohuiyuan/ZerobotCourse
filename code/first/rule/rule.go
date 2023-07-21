package rule

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	engine := zero.New()
	engine.OnFullMatch("牛逼", zero.AdminPermission).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("你很牛逼"))
	})
	engine.OnFullMatch("消息检测", checkRule).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("检测完毕"))
	})
}

func checkRule(ctx *zero.Ctx) bool {
	if zero.OnlyPrivate(ctx) {
		ctx.SendChain(message.Text("这是私聊信息"))
	}
	if zero.OnlyGroup(ctx) {
		ctx.SendChain(message.Text("这是群聊信息"))
	}
	if zero.OnlyToMe(ctx) {
		ctx.SendChain(message.Text("这是@bot的信息"))
	}
	if zero.OnlyGuild(ctx) {
		ctx.SendChain(message.Text("这是频道信息"))
	}
	if zero.OnlyPublic(ctx) {
		ctx.SendChain(message.Text("这是群聊或者频道信息"))
	}
	if zero.AdminPermission(ctx) {
		ctx.SendChain(message.Text("你有管理员权限"))
	}
	if zero.OwnerPermission(ctx) {
		ctx.SendChain(message.Text("你有群主权限"))
	}
	if zero.SuperUserPermission(ctx) {
		ctx.SendChain(message.Text("你有主人权限"))
	}
	return true
}
