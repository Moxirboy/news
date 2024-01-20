package http

import (
	"news/internal/controller/http/v1/dto"
	"news/internal/models"
	"time"
)

func CreateNewsRequestTo(
	req *dto.CreateNewsRequest,
) *models.News {
	return &models.News{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		By: models.By{
			CreatedBy: req.Author,
		},
		At: models.At{
			CreatedAt: time.Now(),
		},
	}
}