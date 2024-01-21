package repo

import( "context"
"news/internal/models"
	"news/pkg/utils"
)

type ICommentRepository interface {
	Create (
		ctx context.Context,
		comment *models.Comment,
		) error
	GetAll(
		ctx context.Context,
		query utils.PaginationQuery,
		) (*models.CommentList,
		error,
		)
	Update(
		ctx context.Context,
		comment *models.Comment,
		) error
	Delete (
		ctx context.Context,
		ID string,
		)error
}