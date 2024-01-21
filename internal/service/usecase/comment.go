package usecase

import (
	"context"
	"news/internal/models"
	"news/internal/service/repo"
	"news/pkg/logger"
	"news/pkg/utils"
)

type commentUseCase struct {
	commentRepo repo.ICommentRepository
	log      logger.Logger
}

func NewCommentUseCase(
	log logger.Logger,
	commentRepo repo.ICommentRepository,
	) ICommentUseCase {
	return &commentUseCase{
		commentRepo: commentRepo,
		log:      log,
		}
	}

	func (uc *commentUseCase) Create(
		ctx context.Context,
		comment *models.Comment) (
		err error) {
		err = uc.commentRepo.Create(ctx, comment)
		if err != nil {
			uc.log.Error("usecase.comment.create Error: ", err)
			return  err
		}
		return nil
		}
	func (uc *commentUseCase) GetAll(
	ctx context.Context,
	query utils.PaginationQuery,
	) (
	comment *models.CommentList,
	err error,
	) {
	comment, err = uc.commentRepo.GetAll(ctx,query)
	if err != nil {
		uc.log.Error("usecase.comment.GetAll Error: ", err)
		return nil, err
	}
	return comment, nil
	}
	func (uc *commentUseCase) Update(
	ctx context.Context,
	comment *models.Comment,
	) error {
	err := uc.commentRepo.Update(ctx, comment)
	if err != nil {
		uc.log.Error("usecase.comment.Update Error: ", err)
		return err
	}
	return nil
	}
	func (uc *commentUseCase) Delete(
	ctx context.Context,
	id string,
	) error {
	err := uc.commentRepo.Delete(ctx, id)
	if err != nil {
		uc.log.Error("usecase.comment.Delete Error: ", err)
		return err
	}
	return nil

	}
