package http

import (
	"news/internal/controller/http/v1/dto"
	"news/internal/models"
	"time"
)


func fromCreateNewsRequest(
	req *dto.NewsRequest,
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
func ToNewsResponse(
	news *models.News,
) dto.NewsResponse {
	return dto.NewsResponse{
		ID:        news.ID,
		Title:     news.Title,
		Content:   news.Content,
		Category:  news.Category,
		Author:    news.CreatedBy,
		CreatedAt: news.CreatedAt.Format("2006-01-02"),
		UpdatedAt: news.UpdatedAt.Format("2006-01-02"),
	}
}
func ToNewsResponseList(
	news *models.NewsList,
	) dto.NewsListResponse{
	lenth:=len(news.News)
	News:=make([]dto.NewsResponse,lenth)
	for _,instance:=range news.News{
		News=append(News,ToNewsResponse(instance))
	}
	return dto.NewsListResponse{
		TotalCount: news.TotalCount,
		TotalPages: news.TotalPages,
		Page: news.Page,
		Size: news.Size,
		HasMore: news.HasMore,
		News: News,
	}
}
