package handler

import (
	"context"

	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/dto"
	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/model"
)

type Service interface {
	Shorten(context.Context, dto.RequestURL) (*model.URL, error)
}
