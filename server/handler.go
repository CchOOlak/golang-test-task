package server

import (
	"fmt"
	"net/http"
	"twitch_chat_analysis/cache"
	"twitch_chat_analysis/messaging"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(200, "worked")
}

func message(c *gin.Context) {
	var json messaging.Message
	err := c.BindJSON(&json)
	if err == nil && json.Sender != "" && json.Receiver != "" {
		err := messaging.Client.Produce(json)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "internal server error",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "message sent successfully",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "sender or receiver is empty",
		})
	}
}

func report(c *gin.Context) {
	sender, receiver := c.Query("sender"), c.Query("receiver")

	if sender == "" || receiver == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "sender or receiver is empty",
		})
	} else {
		key := fmt.Sprintf("%s_%s", sender, receiver)
		messages := cache.Get(key)
		c.JSON(http.StatusOK, gin.H{
			"messages": messages,
		})
	}
}