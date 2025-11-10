package handler

import (
	"context"
	"io"

	"github.com/AugustSerenity/go-contest-L3/l3.4_ImageProcessor/internal/model"
)

type Service interface {
	UploadImage(ctx context.Context, file io.Reader, filename string) (string, error)
	ProcessImage(ctx context.Context, imageID, imagePath string) error
	GetImage(ctx context.Context, id string) (*model.Image, error)
	DeleteImage(ctx context.Context, id string) error
}
