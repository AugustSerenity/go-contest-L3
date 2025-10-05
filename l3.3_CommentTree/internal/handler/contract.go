package handler

import "github.com/AugustSerenity/go-contest-L3/tree/main/l3.3_CommentTree/internal/handler/dto"

type Service interface {
	CreateComment(commentIn dto.CommentRequest) (dto.CommentResponse, error)
}
