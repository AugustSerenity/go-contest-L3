package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/AugustSerenity/go-contest-L3/l3.4_ImageProcessor/internal/service"
	"github.com/segmentio/kafka-go"
	"github.com/wb-go/wbf/retry"
)

type ImageProcessorConsumer struct {
	consumer *kafka.Reader
	service  *service.Service
}

func NewImageProcessorConsumer(brokers []string, topic, groupID string, service *service.Service) *ImageProcessorConsumer {
	return &ImageProcessorConsumer{
		consumer: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: groupID,
		}),
		service: service,
	}
}

func (c *ImageProcessorConsumer) StartConsuming(ctx context.Context, out chan<- kafka.Message, strategy retry.Strategy) error {
	for {
		msg, err := c.FetchWithRetry(ctx, strategy)
		if err != nil {
			log.Printf("Error fetching message: %v", err)
			continue
		}

		var task struct {
			ImageID string `json:"image_id"`
			Path    string `json:"path"`
		}

		if err := json.Unmarshal(msg.Value, &task); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		if err := c.service.ProcessImage(ctx, task.ImageID, task.Path); err != nil {
			log.Printf("Error processing image: %v", err)
			continue
		}

		if err := c.consumer.CommitMessages(ctx, msg); err != nil {
			log.Printf("Error committing message: %v", err)
			continue
		}

		select {
		case out <- msg:
		case <-ctx.Done():
			return nil
		}
	}
}

func (c *ImageProcessorConsumer) FetchWithRetry(ctx context.Context, strategy retry.Strategy) (kafka.Message, error) {
	var msg kafka.Message
	err := retry.Do(func() error {
		m, e := c.consumer.FetchMessage(ctx)
		if e == nil {
			msg = m
		}
		return e
	}, strategy)
	return msg, err
}

func (c *ImageProcessorConsumer) Close() error {
	return c.consumer.Close()
}
