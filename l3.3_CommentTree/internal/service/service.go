package service

import (
	"time"

	"github.com/AugustSerenity/go-contest-L3/tree/main/l3.3_CommentTree/internal/handler/dto"
	"github.com/AugustSerenity/go-contest-L3/tree/main/l3.3_CommentTree/internal/model"
)

type Service struct {
	storage Storage
}

func New(st Storage) *Service {
	return &Service{
		storage: st,
	}
}

func (s *Service) CreateComment(commentIncome dto.CommentRequest) (dto.CommentResponse, error) {
	comment := model.Comment{
		Text:      commentIncome.Text,
		ParentID:  commentIncome.ParentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := s.storage.InsertComment(comment)
	if err != nil {
		return dto.CommentResponse{}, err
	}

	comment.ID = id
	return model.CastModel(comment), nil
}
