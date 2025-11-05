package producer

import (
	"context"

	"github.com/wb-go/wbf/kafka"
)

type KafkaServiceProducer struct {
	producer *kafka.Producer
}

func NewKafkaServiceProducer(producer *kafka.Producer) *KafkaServiceProducer {
	return &KafkaServiceProducer{producer: producer}
}

func (k *KafkaServiceProducer) Send(ctx context.Context, key, value []byte) error {
	return k.producer.Send(ctx, key, value)
}
