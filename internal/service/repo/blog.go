

package repo

import (
	"context"
	"news/internal/models"
	"news/pkg/utils"
)

type IBlogRepository interface {
	Create(
		ctx context.Context,
		blog *models.Blog,
		) (blogID string, err error)
	GetByID(
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