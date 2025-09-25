package producer

import "github.com/AugustSerenity/go-contest-L3/l3.1/internal/model"

type ProducerService interface {
	Publish(model.Notification) error
}
