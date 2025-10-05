package dto

import "github.com/AugustSerenity/go-contest-L3/tree/main/l3.3_CommentTree/internal/model"

func CastModel(c model.Comment) CommentResponse {
	return CommentResponse{
		ID:       c.ID,
		Text:     c.Text,
		ParentID: c.ParentID,
	}
}
