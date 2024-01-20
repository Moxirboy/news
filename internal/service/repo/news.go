package repo

import (
	"context"
	"news/internal/models"
	"news/pkg/utils"
)

type INewsRepository interface {
	Create(
		ctx context.Context,
		news *models.News,
	) (newsID string, err error)
	GetByID(
		ctx context.Context,
		ID string) (
		*models.News, error,
		)
	GetAll(
		ctx context.Context,
		query utils.PaginationQuery,
		) (
		*models.NewsList,
		error,
		)
	Update(
		ctx context.Context,
		news *models.News,
		) error
	Delete(
		ctx context.Context,
		id string,
		) error
}