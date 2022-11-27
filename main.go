package main

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
)

func main() {
	zero.OnCommand("八股文", HandleBagu)

	zero.RunAndBlock(&zero.Config{
		NickName:      []string{"bot"},
		CommandPrefix: "/",
		SuperUsers:    []int64{10000},
		Driver: []zero.Driver{
			driver.NewWebSocketClient("ws://42.192.138.109:7778/", ""),
		},
	}, nil)
}
