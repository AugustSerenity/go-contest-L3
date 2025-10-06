package service

import (
	"encoding/json"
	"time"

	"github.com/AugustSerenity/go-contest-L3/tree/main/l3.3_CommentTree/internal/handler/dto"
	"github.com/AugustSerenity/go-contest-L3/tree/main/l3.3_CommentTree/internal/model"
	"github.com/wb-go/wbf/zlog"
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

func (s *Service) SearchComments(q string, page, limit int) ([]dto.CommentResponse, error) {
	comments, err := s.storage.SearchComments(q, page, limit)
	if err != nil {
		return nil, err
	}

	tree := buildTree(comments)
	return tree, nil
}

func (s *Service) GetAllComments(idComment string) ([]dto.CommentResponse, error) {
	comments, err := s.storage.GetTree(idComment)
	if err != nil {
		return nil, err
	}
	tree := buildTree(comments)

	// DEBUG
	if debugJson, err := json.MarshalIndent(tree, "", "  "); err == nil {
		zlog.Logger.Info().Msg("Full comment tree:\n" + string(debugJson))
	}

	return tree, nil
}

func buildTree(comments []model.Comment) []dto.CommentResponse {
	idToComment := make(map[int64]*dto.CommentResponse)

	for _, c := range comments {
		copied := model.CastModel(c)
		idToComment[c.ID] = &copied
	}

	zlog.Logger.Info().Int("total_nodes", len(idToComment)).Msg("Starting buildTree")

	var roots []*dto.CommentResponse
	for _, c := range comments {
		current := idToComment[c.ID]

		if c.ParentID != nil {
			parent, ok := idToComment[*c.ParentID]
			if ok {
				zlog.Logger.Debug().
					Int64("child_id", c.ID).
					Int64("parent_id", *c.ParentID).
					Msg("Attaching child to parent")

				parent.Children = append(parent.Children, *current)
			} else {
				zlog.Logger.Warn().
					Int64("child_id", c.ID).
					Int64("parent_id", *c.ParentID).
					Msg("Parent not found — skipping child")
			}
		} else {
			roots = append(roots, current)
		}
	}

	result := make([]dto.CommentResponse, len(roots))
	for i, r := range roots {
		result[i] = *r
	}

	zlog.Logger.Info().Int("root_nodes", len(result)).Msg("Finished buildTree")
	return result
}

func (s *Service) DeleteComment(id string) error {
	err := s.storage.DeleteCommentByID(id)
	if err != nil {
		return err
	}

	return nil
}
