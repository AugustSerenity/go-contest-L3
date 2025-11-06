package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/AugustSerenity/go-contest-L3/l3.4_ImageProcessor/internal/kafka/producer"
	"github.com/AugustSerenity/go-contest-L3/l3.4_ImageProcessor/internal/model"
)

type Service struct {
	storage     Storage
	producer    producer.ServiceProducer
	storagePath string
}

func New(st Storage, sp producer.ServiceProducer, storagePath string) *Service {
	return &Service{
		storage:     st,
		producer:    sp,
		storagePath: storagePath,
	}
}

func (s *Service) UploadImage(ctx context.Context, file io.Reader, filename string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(filename + time.Now().String()))
	imageID := hex.EncodeToString(hash.Sum(nil))[:16]

	originalPath := filepath.Join(s.storagePath, "originals", imageID+filepath.Ext(filename))
	err := s.storage.SaveFile(file, originalPath)
	if err != nil {
		return "", err
	}

	image := &model.Image{
		ID:           imageID,
		OriginalPath: originalPath,
		Status:       model.StatusPending,
		CreatedAt:    time.Now(),
	}

	if err := s.storage.Create(ctx, image); err != nil {
		os.Remove(originalPath)
		return "", err
	}

	message := map[string]string{
		"image_id": imageID,
		"path":     originalPath,
	}
	messageBytes, _ := json.Marshal(message)

	if err := s.producer.Send(ctx, []byte(imageID), messageBytes); err != nil {
		return "", err
	}

	return imageID, nil

}

func (s *Service) ProcessImage(ctx context.Context, imageID, imagePath string) error {
	image, err := s.storage.GetByID(ctx, imageID)
	if err != nil {
		return fmt.Errorf("failed to get image: %w", err)
	}

	image.Status = model.StatusProcessing
	image.ProcessedAt = &time.Time{}
	if err := s.storage.UpdateStatus(ctx, image); err != nil {
		return fmt.Errorf("failed to update image status: %w", err)
	}

	image.Status = model.StatusCompleted
	if err := s.storage.UpdateStatus(ctx, image); err != nil {
		return fmt.Errorf("failed to update image status: %w", err)
	}

	return nil
}
