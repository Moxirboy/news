package postgres

import (
	"context"
	"database/sql"
	"news/internal/models"
	"news/internal/service/repo"
	"news/pkg/logger"
)

type newsRepository struct {
	db *sql.DB
	logger logger.Logger
}
func NewNewsRepo(db *sql.DB,logger logger.Logger) repo.INewsRepository{
	return &newsRepository{
		db:db,
		logger:logger,
	}
}
func (r *newsRepository) Create(
	ctx context.Context,
	news *models.News,
	) (newsID string, err error){
	err = r.db.QueryRowContext(
		ctx,
		NewsCreate,
		news.Title,
		news.Content,
		news.CreatedBy,
		news.Category,
		news.CreatedAt,
	).Scan(&newsID)
	if err != nil {
		return "", err
	}
	return newsID, nil
}
	