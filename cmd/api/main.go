package main

import (
	"twitch_chat_analysis/cache"
	"twitch_chat_analysis/messaging"
	"twitch_chat_analysis/server"
)

func main() {
	cache.Init()
	messaging.Init()
	r := server.InitServer()
	r.Run()
}
