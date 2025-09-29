package handler

import (
	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/dto"
	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/model"
	"github.com/wb-go/wbf/ginext"
)

type Service interface {
	Shorten(*ginext.Context, dto.RequestURL) (model.URL, error)
}
