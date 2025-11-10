package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/AugustSerenity/go-contest-L3/l3.4_ImageProcessor/internal/kafka/producer"
	"github.com/AugustSerenity/go-contest-L3/l3.4_ImageProcessor/internal/model"
	"github.com/disintegration/imaging"
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
	imageData, err := s.storage.GetByID(ctx, imageID)
	if err != nil {
		return fmt.Errorf("failed to get image: %w", err)
	}

	imageData.Status = model.StatusProcessing
	_ = s.storage.UpdateStatus(ctx, imageData)

	img, err := imaging.Open(imagePath)
	if err != nil {
		imageData.Status = model.StatusFailed
		_ = s.storage.UpdateStatus(ctx, imageData)
		return fmt.Errorf("open image: %w", err)
	}

	processedDir := filepath.Join(s.storagePath, "processed")
	thumbnailsDir := filepath.Join(s.storagePath, "thumbnails")
	os.MkdirAll(processedDir, os.ModePerm)
	os.MkdirAll(thumbnailsDir, os.ModePerm)

	resized := imaging.Resize(img, 1024, 0, imaging.Lanczos)
	resizedPath := filepath.Join(processedDir, imageID+"_resized.jpg")
	_ = imaging.Save(resized, resizedPath)

	thumb := imaging.Thumbnail(img, 200, 200, imaging.Lanczos)
	thumbPath := filepath.Join(thumbnailsDir, imageID+"_thumb.jpg")
	_ = imaging.Save(thumb, thumbPath)

	wmPath := filepath.Join(s.storagePath, "watermark.png")
	final := resized
	if wm, err := imaging.Open(wmPath); err == nil {
		final = imaging.Overlay(resized, wm, image.Pt(20, 20), 0.4)
	}
	watermarkPath := filepath.Join(processedDir, imageID+"_watermarked.jpg")
	_ = imaging.Save(final, watermarkPath)

	now := time.Now()
	imageData.Status = model.StatusCompleted
	imageData.ProcessedAt = &now
	imageData.ResizedPath = resizedPath
	imageData.ThumbPath = thumbPath
	imageData.WatermarkPath = watermarkPath

	return s.storage.UpdateStatus(ctx, imageData)
}

func (s *Service) GetImage(ctx context.Context, id string) (*model.Image, error) {
	return s.storage.GetByID(ctx, id)
}

func (s *Service) DeleteImage(ctx context.Context, id string) error {
	image, err := s.storage.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if err := s.storage.Delete(ctx, id); err != nil {
		return err
	}
	os.Remove(image.OriginalPath)
	if image.ProcessedAt != nil {
		processedPath := filepath.Join(s.storagePath, "processed", image.ID+filepath.Ext(image.OriginalPath))
		os.Remove(processedPath)
	}
	return nil
}
