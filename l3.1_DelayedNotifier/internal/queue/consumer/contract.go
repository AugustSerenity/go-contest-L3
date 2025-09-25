package consumer

import "github.com/AugustSerenity/go-contest-L3/l3.1/internal/model"

type ConsumerService interface {
	ProcessNotification(model.Notification)
}
