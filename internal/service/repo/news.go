package repo

import (
	"context"
	"news/internal/models"
)

type INewsRepository interface {
	Create(
		ctx context.Context,
		news *models.News,
	) (newsID string, err error)
	
}