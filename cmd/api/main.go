package main

import (
	"twitch_chat_analysis/server"
)

func main() {
	r := server.InitServer()
	r.Run()
}
