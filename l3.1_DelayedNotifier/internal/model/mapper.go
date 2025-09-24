package model

import (
	"time"

	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/handler/dto"
	"github.com/google/uuid"
)

func CastToNotification(request dto.NotificationRequest) Notification {
	return Notification{
		ID:        uuid.New().String(),
		Message:   request.Message,
		SendAt:    request.SendAt,
		CreatedAt: time.Now().UTC(),
		Status:    "scheduled",
	}
}
