package api_server

import (
  "context"
	kafka "github.com/segmentio/kafka-go"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Message struct {
	ID int `json:"id" binding:"required"`
	MSG string `json:"msg" binding:"required"`
}

func MsgHandler(c *gin.Context, w *kafka.Writer) {
	var msg_struct Message

	if err := c.BindJSON(&msg_struct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
      "Status": err.Error(),
    })
		return
	}
	data, _ := json.Marshal(msg_struct)

	// Produce messages to topic
	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   nil,
			Value: data,
		},
	)

	c.JSON(http.StatusOK, gin.H{
		"Id": msg_struct.ID,
		"Msg": msg_struct.MSG,
    "Status": "Published",
	})
}
