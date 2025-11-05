package service

import (
	"context"
	"io"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) UploadImage(ctx context.Context, file io.Reader, filename string) (string, error) {
	
}
