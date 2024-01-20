package postgres

import (
	"context"
	"database/sql"
	"news/internal/models"
	"news/internal/service/repo"
	"news/pkg/logger"
	"news/pkg/utils"
	"time"
)

type newsRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewNewsRepo(db *sql.DB, logger logger.Logger) repo.INewsRepository {
	return &newsRepository{
		db:     db,
		logger: logger,
	}
}
func (r *newsRepository) Create(
	ctx context.Context,
	news *models.News,
) (newsID string, err error) {
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
func (r *newsRepository) GetByID(
	ctx context.Context,
	ID string) (
	*models.News, error,
) {
	news := models.News{}
	err := r.db.QueryRowContext(
		ctx,
		NewsGetById,
		ID,
	).Scan(
		&news.ID,
		&news.Title,
		&news.Content,
		&news.CreatedBy,
		&news.Category,
		&news.CreatedAt,
	)
	if err != nil {
		r.logger.Error("repo.news.GetById Error:", err)
		return nil, err
	}
	return &news, nil
}
func (r *newsRepository) GetAll(
	ctx context.Context,
	query utils.PaginationQuery,
) (
	*models.NewsList,
	error,
) {
	count := 0
	if err := r.db.QueryRowContext(ctx,
		NewsCount).Scan(
		&count); err != nil {
		r.logger.Error("repo.news.GetAll Error:", err)
		return nil, err
	}
	if count == 0 {
		return &models.NewsList{
			TotalCount: count,
			TotalPages: utils.GetTotalPages(count, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), count, query.GetSize()),
			News:       make([]*models.News, 0),
		}, nil
	}

	rows, err := r.db.QueryContext(
		ctx,
		NewsGetAll,
		query.GetOffset(),
		query.GetLimit(),
	)
	if err != nil {
		r.logger.Error("repo.news.GetAll Error: ", err)
		return nil, err
	}
	defer rows.Close()
	news := make([]*models.News, 0, query.GetSize())
	for rows.Next() {
		new := &models.News{}
		rows.Scan(
			&new.ID,
			&new.Title,
			&new.Content,
			&new.CreatedBy,
			&new.Category,
			&new.CreatedAt,
		)
		news = append(news, new)
	}
	return &models.NewsList{
		TotalCount: count,
		TotalPages: utils.GetTotalPages(count, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), count, query.GetSize()),
		News:       news,
	}, nil
}
func (r *newsRepository) Update(
	ctx context.Context,
	news *models.News,
) error {
	tx, err := r.db.BeginTx(
		context.Background(),
		&sql.TxOptions{
			Isolation: sql.LevelSerializable,
		},
	)
	if err != nil {
		r.logger.Error("repo.news.update error while transaction begin:", err)
		return err
	}
	res, execErr := tx.ExecContext(
		ctx,
		NewsUpdate,
		news.Title,
		news.Content,
		news.CreatedBy,
		news.Category,
		news.ID,
	)
	if execErr != nil {
		r.logger.Error(
			"repo.box.udpate error while insert box:",
			execErr.Error(),
		)
		_ = tx.Rollback()
		return execErr
	}

	if count, _ := res.RowsAffected(); count == 0 {
		_ = tx.Rollback()
		return sql.ErrNoRows
	}
	return nil
}
func (r *newsRepository) Delete(
	ctx context.Context,
	id string,
) error {
	tx, err := r.db.BeginTx(
		context.Background(),
		&sql.TxOptions{
			Isolation: sql.LevelSerializable,
		},
	)
	if err != nil {
		r.logger.Error("repo.news.update error while transaction begin:", err)
		return err
	}

	res, execErr := tx.ExecContext(
		ctx,
		NewsDelete,
		id,
		time.Now().Format("2006-01-02"),
	)
	if execErr != nil {
		r.logger.Error(
			"repo.box.udpate error while insert box:",
			execErr.Error(),
		)
		_ = tx.Rollback()
		return execErr
	}

	if count, _ := res.RowsAffected(); count == 0 {
		_ = tx.Rollback()
		return sql.ErrNoRows
	}
	
	return nil

}
