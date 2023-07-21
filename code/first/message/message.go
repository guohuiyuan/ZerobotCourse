package message

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	engine := zero.New()
	engine.OnFullMatch("文本").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("文本"))
	})
	engine.OnFullMatch("艾特").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.At(ctx.Event.UserID), message.Text("艾特你"))
	})
	engine.OnFullMatch("回复").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text("回复你"))
	})
	engine.OnFullMatch("图片").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image("https://gitcode.net/anto_july/imagematerials/-/blob/main/need/0.png"))
	})
	engine.OnFullMatch("语音").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Record("https://vtbkeyboard.moe/api/audio/672328094/%E7%8C%AB%E7%8C%AB%E8%B7%9F%E9%BC%A0%E9%BC%A0%E4%B8%8D%E5%86%B2%E7%AA%81.mp3"))
	})
	engine.OnFullMatch("表情").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Face(1))
	})
	engine.OnFullMatch("音乐").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Music("163", 28949129))
	})
	engine.OnFullMatch("tts").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.TTS("你好"))
	})
	engine.OnFullMatch("戳一戳").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Poke(ctx.Event.UserID))
	})
}
