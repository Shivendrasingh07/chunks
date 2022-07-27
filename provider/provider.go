package provider

import (
	"context"
	"example.com/m/models"
	"mime/multipart"
	"time"
)

type StorageProvider interface {
	UploadV3(ctx context.Context, chunk *models.Chunk, bucketName string, file multipart.File, filePath string) error
	GetSharableURL(bucketName, fileName string, expireTimeInHours time.Duration) (string, error)
}
