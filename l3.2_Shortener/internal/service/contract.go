package service

import (
	"context"

	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/model"
)

type Storage interface {
	SaveLink(context.Context, *model.URL) error
	ExistsByShortCode(context.Context, string) (bool, error)
}
