package service

import (
	"context"

	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/model"
)

type Storage interface {
	SaveLink(context.Context, *model.URL) error
	ExistsByShortCode(context.Context, string) (bool, error)
	GetOriginalURL(context.Context, string) (string, error)
	GetLinkIDByShortURL(ctx context.Context, shortURL string) (int, error)
	GetClicksByLinkID(ctx context.Context, linkID int) ([]model.Click, error)
	InsertClick(ctx context.Context, linkID int, userAgent string) error
}
