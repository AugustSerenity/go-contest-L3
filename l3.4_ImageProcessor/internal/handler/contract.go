package handler

import (
	"context"
	"io"
)

type Service interface {
	UploadImage(ctx context.Context, file io.Reader, filename string) (string, error)
}
