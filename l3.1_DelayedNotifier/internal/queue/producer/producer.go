package producer

import (
	"encoding/json"
	"time"

	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/model"
	"github.com/wb-go/wbf/rabbitmq"
	"github.com/wb-go/wbf/retry"
)

type Producer struct {
	publisher *rabbitmq.Publisher
	queueName string
}

func New(p *rabbitmq.Publisher, queueName string) *Producer {
	return &Producer{
		publisher: p,
		queueName: queueName,
	}
}

func (p *Producer) Publish(n model.Notification) error {
	body, err := json.Marshal(&n)
	if err != nil {
		return err
	}

	return p.publisher.PublishWithRetry(
		body,
		p.queueName,
		"application/json",
		retry.Strategy{
			Attempts: 3,
			Delay:    300 * time.Millisecond,
			Backoff:  2.0,
		},
	)
}
