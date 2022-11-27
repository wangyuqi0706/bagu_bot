package main

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func HandleBagu(ctx *zero.Ctx) bool {
	messageID := ctx.Send(ctx.Send(message.Text("测试文本").String()))
	return messageID.ID() == 0
}
