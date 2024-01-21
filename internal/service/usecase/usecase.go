package usecase

import (
	"context"
	"news/internal/models"
	"news/pkg/utils"
)
type INewsUsecase interface{
	Create(
		ctx context.Context,
		news *models.News,
		) (newsID string, err error)
	GetOneByID(
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
type IBlogUsecase interface{
	Create(
		ctx context.Context,
		blog *models.Blog,
		) (blogID string, err error)
	GetOneByID(
		ctx context.Context,
		ID string) (
		*models.Blog, error,
		)
	GetAll(
		ctx context.Context,
		query utils.PaginationQuery,
		) (
		*models.BlogList,
		error,
		)
	Update(
		ctx context.Context,
		blog *models.Blog,
		) error
	Delete(
		ctx context.Context,
		id string,
		) error
}