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
