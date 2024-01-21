package usecase

import (
	"context"
	"news/internal/models"
	"news/internal/service/repo"
	"news/pkg/logger"
	"news/pkg/utils"
)

type blogUseCase struct {
	blogRepo repo.IBlogRepository
	log      logger.Logger
}

func NewblogUseCase(
	log logger.Logger,
	blogRepo repo.IBlogRepository,
	) IBlogUsecase {
	return &blogUseCase{
		blogRepo: blogRepo,
		log:      log,
		}
	}
	func (uc *blogUseCase) Create(
		ctx context.Context,
		blog *models.Blog,
		) (blogID string, err error) {
	blogID, err = uc.blogRepo.Create(ctx, blog)
	if err != nil {
		uc.log.Error("usecase.blog.create Error: ", err)
		return "", err
	}
	return blogID, nil
	}
	func (uc *blogUseCase)GetOneByID(
		ctx context.Context,
		ID string) (
		*models.Blog, error,
		) {
	blog, err := uc.blogRepo.GetByID(ctx, ID)
	if err != nil {
		uc.log.Error("usecase.blog.GetOneByID Error: ", err)
		return nil, err
	}
	return blog, nil
	}
	func (uc *blogUseCase) GetAll(
		ctx context.Context,
		query utils.PaginationQuery,
		) (
		*models.BlogList,
		error,
		) {
	blog, err := uc.blogRepo.GetAll(ctx,query)
	if err != nil {
		uc.log.Error("usecase.blog.GetAll Error: ", err)
		return nil, err
	}
	return blog, nil
	}
	func (uc *blogUseCase) Update(
		ctx context.Context,
		blog *models.Blog,
		) error{
	err := uc.blogRepo.Update(ctx, blog)
	if err != nil {
		uc.log.Error("usecase.blog.Update Error: ", err)
		return err
	}
	return nil
	}
	func (uc *blogUseCase) Delete(
		ctx context.Context,
		id string,
		) error {
	err := uc.blogRepo.Delete(ctx, id)
	if err != nil {
		uc.log.Error("usecase.blog.Delete Error: ", err)
		return err
	}
	return nil

	}
