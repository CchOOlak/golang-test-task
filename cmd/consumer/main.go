package main

import (
	"twitch_chat_analysis/cache"
	"twitch_chat_analysis/messaging"
)

func main() {
	messaging.Init()
	messaging.Client.Consume(
		cache.ProcessMessage,
	)
}