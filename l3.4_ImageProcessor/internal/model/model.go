package model

import "time"

type ImageStatus string

const (
	StatusPending    ImageStatus = "pending"
	StatusProcessing ImageStatus = "processing"
	StatusCompleted  ImageStatus = "completed"
	StatusFailed     ImageStatus = "failed"
)

type Image struct {
	ID            string      `json:"id"`
	OriginalPath  string      `json:"original_path"`
	ResizedPath   string      `json:"resized_path,omitempty"`
	ThumbPath     string      `json:"thumb_path,omitempty"`
	WatermarkPath string      `json:"watermark_path,omitempty"`
	Status        ImageStatus `json:"status"`
	CreatedAt     time.Time   `json:"created_at"`
	ProcessedAt   *time.Time  `json:"processed_at"`
}
