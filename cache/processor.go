package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"twitch_chat_analysis/messaging"
)

func ProcessMessage(body []byte) {
	msg := messaging.Message{}
	json.Unmarshal(body, &msg)

	key := fmt.Sprintf("%s_%s", msg.Sender, msg.Receiver)
	Client.Add(key, msg.Message)
	log.Default().Printf("new message %s added to %s\n", msg.Message, key)
}