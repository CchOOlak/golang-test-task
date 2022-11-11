package cache

import (
	"encoding/json"
	"fmt"
	"twitch_chat_analysis/messaging"
)

func ProcessMessage(body []byte) {
	msg := messaging.Message{}
	json.Unmarshal(body, &msg)

	key := fmt.Sprintf("%s_%s", msg.Sender, msg.Receiver)
	Add(key, msg.Message)
}