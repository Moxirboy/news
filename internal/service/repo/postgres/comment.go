package postgres

import (
	"database/sql"
	"context"
	"news/internal/models"
	"news/pkg/utils"
	"news/internal/service/repo"
	"news/pkg/logger"
	
)

type commentRepository struct{
	db *sql.DB
	log logger.Logger
}
func NewCommentRepository(
	sql *sql.DB,
	log logger.Logger,
	) repo.ICommentRepository{
	return &commentRepository{
		sql,
		log,
	}
}
func (r commentRepository)Create (
	ctx context.Context,
	comment *models.Comment,
	) error{
	_,err := r.db.ExecContext(
		ctx,
		createComment,
		comment.BlogId,
		comment.Body,
		)
	if err != nil {
		return err
	}
	return  nil
}
func (r commentRepository)GetAll(
	ctx context.Context,
	query utils.PaginationQuery,
	) (*models.CommentList,
	error,
	){
	count := 0
	if err := r.db.QueryRowContext(ctx,
		countComment).Scan(
			&count); err != nil {
		r.log.Error("repo.comment.GetAll Error:", err)
		return nil, err
			}
			if count == 0 {
				return &models.CommentList{
					TotalCount: count,
					TotalPages: utils.GetTotalPages(count, query.GetSize()),
					Page:       query.GetPage(),
					Size:       query.GetSize(),
					HasMore:    utils.GetHasMore(query.GetPage(), count, query.GetSize()),
					Comment:       make([]*models.Comment, 0),
					}, nil
			}

			rows, err := r.db.QueryContext(
				ctx,
				getAllComment,
				query.GetOffset(),
				query.GetLimit(),
				)
			if err != nil {
				r.log.Error("repo.comment.GetAll Error: ", err)
				return nil, err
			}
			defer rows.Close()
			commentSlice:= make([]*models.Comment, 0, query.GetSize())
			for rows.Next() {
				comment := &models.Comment{}
				rows.Scan(
					&comment.ID,
	&comment.BlogId,
	&comment.Body,
					)
				commentSlice = append(commentSlice, comment)
			}
			return &models.CommentList{
				TotalCount: count,
				TotalPages: utils.GetTotalPages(count, query.GetSize()),
				Page:       query.GetPage(),
				Size:       query.GetSize(),
				HasMore:    utils.GetHasMore(query.GetPage(), count, query.GetSize()),
				Comment:       commentSlice,
				}, nil
}
func (r commentRepository)Update(
		ctx context.Context,
		comment *models.Comment,
		) error{
	
	tx, err := r.db.BeginTx(
		context.Background(),
		&sql.TxOptions{
			Isolation: sql.LevelSerializable,
			},
			)
	if err != nil {
		r.log.Error("repo.comment.update error while transaction begin:", err)
		return err
	}
	res, execErr := tx.ExecContext(
		ctx,
		updateComment,
		comment.Body,
		)
	if execErr != nil {
		r.log.Error(
			"repo.comment.udpate error while insert box:",
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
func (r commentRepository)Delete (
	ctx context.Context,
	ID string,
	)error{
	tx, err := r.db.BeginTx(
		context.Background(),
		&sql.TxOptions{
			Isolation: sql.LevelSerializable,
			},
			)
	if err != nil {
		r.log.Error("repo.comment.update error while transaction begin:", err)
		return err
	}

	res, execErr := tx.ExecContext(
		ctx,
		deleteComment,
		ID,
		)
	if execErr != nil {
		r.log.Error(
			"repo.comment.udpate error while insert box:",
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

