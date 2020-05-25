package api_server

import (
	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
	"github.com/gouravtulsani/golang-http-server/api_server/http_handler"
)

func Init() {
  BROKER_URL := []string{"192.168.1.9:9092"}
  KAFKA_TOPIC := "test-insert"

	// create kafka producer
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  BROKER_URL,
		Topic:    KAFKA_TOPIC,
		Balancer: &kafka.LeastBytes{},
	})

	defer w.Close()

	r := gin.Default()
	r.POST("/msg", func(c *gin.Context) {http_handler.MsgHandler(c, w)})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

