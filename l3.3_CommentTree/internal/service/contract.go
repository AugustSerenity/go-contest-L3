package service

import "github.com/AugustSerenity/go-contest-L3/tree/main/l3.3_CommentTree/internal/model"

type Storage interface {
	InsertComment(comment model.Comment) (int64, error)
}
