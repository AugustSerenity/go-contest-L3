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

func (s *Service) GetAllComments(idComment string) ([]dto.CommentResponse, error) {
	comments, err := s.storage.GetTree(idComment)
	if err != nil {
		return nil, err
	}

	tree := buildTree(comments)
	return tree, nil
}

func buildTree(comments []model.Comment) []dto.CommentResponse {
	idToComment := make(map[int64]*dto.CommentResponse)
	var roots []dto.CommentResponse

	for _, c := range comments {
		dtoComment := model.CastModel(c)
		dtoComment.Children = []dto.CommentResponse{}

		idToComment[c.ID] = &dtoComment
	}

	for _, c := range comments {
		current := idToComment[c.ID]
		if c.ParentID != nil {
			parent := idToComment[*c.ParentID]
			parent.Children = append(parent.Children, *current)
		} else {
			roots = append(roots, *current)
		}
	}

	return roots
}

func (s *Service) DeleteComment(id string) error {
	err := s.storage.DeleteCommentByID(id)
	if err != nil {
		return err
	}

	return nil
}
