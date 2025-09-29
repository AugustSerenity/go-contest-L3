package service

import (
	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/dto"
	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/model"
	"github.com/wb-go/wbf/ginext"
)

type Service struct {
	storage Storage
}

func New(st Storage) *Service {
	return &Service{
		storage: st,
	}
}

func (s *Service) Shorten(c ginext.Context, urlRequest dto.RequestURL) (model.URL, error) {

}
