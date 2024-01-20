package usecase

import (
	"context"
	"news/internal/models"
	"news/internal/service/repo"
	"news/pkg/logger"
	"news/pkg/utils"
)

type newsUseCase struct {
	newsRepo repo.INewsRepository
	log      logger.Logger
}

func NewNewsUseCase(
	log logger.Logger,
	newsRepo repo.INewsRepository,
) INewsUsecase {
	return &newsUseCase{
		newsRepo: newsRepo,
		log:      log,
	}
}
func (uc *newsUseCase) Create(
	ctx context.Context,
	news *models.News) (
	newsID string,
	err error) {
	newsID, err = uc.newsRepo.Create(ctx, news)
	if err != nil {
		uc.log.Error("usecase.news.create Error: ", err)
		return "", err
	}
	return newsID, nil
}
func (uc *newsUseCase) GetOneByID(
	ctx context.Context,
	ID string,
) (
	*models.News, error,
) {
	news, err := uc.newsRepo.GetByID(ctx, ID)
	if err != nil {
		uc.log.Error("usecase.news.GetOneByID Error: ", err)
		return nil, err
	}
	return news, nil
}
func (uc *newsUseCase) GetAll(
	ctx context.Context,
	query utils.PaginationQuery,
) (
news *models.NewsList,
err error,
) {
	news, err = uc.newsRepo.GetAll(ctx,query)
	if err != nil {
		uc.log.Error("usecase.news.GetAll Error: ", err)
		return nil, err
	}
	return news, nil
}
func (uc *newsUseCase) Update(
	ctx context.Context,
	news *models.News,
) error {
	err := uc.newsRepo.Update(ctx, news)
	if err != nil {
		uc.log.Error("usecase.new.Update Error: ", err)
		return err
	}
	return nil
}
func (uc *newsUseCase) Delete(
	ctx context.Context,
	id string,
) error {
	err := uc.newsRepo.Delete(ctx, id)
	if err != nil {
		uc.log.Error("usecase.new.Delete Error: ", err)
		return err
	}
	return nil

}
