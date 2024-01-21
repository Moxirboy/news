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

type blogRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewBlogRepo(db *sql.DB, logger logger.Logger) repo.IBlogRepository {
	return &blogRepository{
		db:     db,
		logger: logger,
		}
}
func (r *blogRepository) Create(
	ctx context.Context,
	blog *models.Blog,
	) (blogID string, err error) {
	err = r.db.QueryRowContext(
		ctx,
		BlogCreate,
		blog.Title,
		blog.Content,
		blog.CreatedBy,
		blog.CreatedAt,
		).Scan(&blogID)
	if err != nil {
		return "", err
	}
	return blogID, nil
	}
	func (r *blogRepository) GetByID(
	ctx context.Context,
	ID string) (
	*models.Blog, error,
	) {
	blog:= models.Blog{}
	err := r.db.QueryRowContext(
		ctx,
		BlogGetById,
		ID,
		).Scan(
			&blog.ID,
			&blog.Title,
			&blog.Content,
			&blog.CreatedBy,
			&blog.CreatedAt,
			)
	if err != nil {
		r.logger.Error("repo.blog.GetById Error:", err)
		return nil, err
	}
	return &blog, nil
	}
	func (r *blogRepository) GetAll(
	ctx context.Context,
	query utils.PaginationQuery,
	) (
	*models.BlogList,
	error,
	) {
	count := 0
	if err := r.db.QueryRowContext(ctx,
		BlogCount).Scan(
			&count); err != nil {
		r.logger.Error("repo.blog.GetAll Error:", err)
		return nil, err
			}
			if count == 0 {
				return &models.BlogList{
					TotalCount: count,
					TotalPages: utils.GetTotalPages(count, query.GetSize()),
					Page:       query.GetPage(),
					Size:       query.GetSize(),
					HasMore:    utils.GetHasMore(query.GetPage(), count, query.GetSize()),
					Blog:       make([]*models.Blog, 0),
					}, nil
			}

			rows, err := r.db.QueryContext(
				ctx,
				BlogGetAll,
				query.GetOffset(),
				query.GetLimit(),
				)
			if err != nil {
				r.logger.Error("repo.blog.GetAll Error: ", err)
				return nil, err
			}
			defer rows.Close()
			blogSlice:= make([]*models.Blog, 0, query.GetSize())
			for rows.Next() {
				blog := &models.Blog{}
				rows.Scan(
					&blog.ID,
					&blog.Title,
					&blog.Content,
					&blog.CreatedBy,
					&blog.CreatedAt,
					)
				blogSlice = append(blogSlice, blog)
			}
			return &models.BlogList{
				TotalCount: count,
				TotalPages: utils.GetTotalPages(count, query.GetSize()),
				Page:       query.GetPage(),
				Size:       query.GetSize(),
				HasMore:    utils.GetHasMore(query.GetPage(), count, query.GetSize()),
				Blog:       blogSlice,
				}, nil
	}
	func (r *blogRepository) Update(
	ctx context.Context,
	blog *models.Blog,
	) error {
	tx, err := r.db.BeginTx(
		context.Background(),
		&sql.TxOptions{
			Isolation: sql.LevelSerializable,
			},
			)
	if err != nil {
		r.logger.Error("repo.blog.update error while transaction begin:", err)
		return err
	}
	res, execErr := tx.ExecContext(
		ctx,
		NewsUpdate,
		blog.Title,
		blog.Content,
		blog.CreatedBy,
		blog.ID,
		)
	if execErr != nil {
		r.logger.Error(
			"repo.blog.udpate error while insert box:",
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
	func (r *blogRepository) Delete(
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
		r.logger.Error("repo.blog.update error while transaction begin:", err)
		return err
	}

	res, execErr := tx.ExecContext(
		ctx,
		BlogDelete,
		id,
		time.Now().Format("2006-01-02"),
		)
	if execErr != nil {
		r.logger.Error(
			"repo.blog.udpate error while insert box:",
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
