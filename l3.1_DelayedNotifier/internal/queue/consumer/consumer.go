package consumer

import (
	"encoding/json"
	"log"

	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/model"
	"github.com/wb-go/wbf/rabbitmq"
)

type Consumer struct {
	consumer *rabbitmq.Consumer
	service  ConsumerService
}

func New(c *rabbitmq.Consumer, s ConsumerService) *Consumer {
	return &Consumer{
		consumer: c,
		service:  s,
	}
}

func (c *Consumer) Start() {
	msgChan := make(chan []byte)

	go func() {
		for msg := range msgChan {
			var n model.Notification
			if err := json.Unmarshal(msg, &n); err != nil {
				log.Printf("failed to decode message: %v", err)
				continue
			}

			c.service.ProcessNotification(n)
		}
	}()

	go func() {
		if err := c.consumer.Consume(msgChan); err != nil {
			log.Fatalf("failed to consume: %v", err)
		}
	}()
}
