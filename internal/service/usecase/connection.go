package usecase

import (
	"database/sql"
	"news/internal/service/repo/postgres"
	"news/pkg/logger"
)

type IUseCase interface {
	NewsUseCase() INewsUsecase
	BlogUseCase() IBlogUsecase
}
type UseCase struct {
	connections map[string]interface{}
}

const (
	_newsUseCase = "news_use_case"
	_blogUseCase = "blog_use_case"
)

func New(
	db *sql.DB,
	log logger.Logger,
) IUseCase {
	var connections = make(map[string]interface{})
	connections[_newsUseCase] =NewNewsUseCase(
		log,
		postgres.NewNewsRepo(db,log),
		)
	connections[_blogUseCase] =NewblogUseCase(
		log,
		postgres.NewBlogRepo(db,log),
		)
	return &UseCase{
		connections: connections,
	}
}
func (c *UseCase) NewsUseCase() INewsUsecase {
	return c.connections[_newsUseCase].(INewsUsecase)
}
func (c *UseCase) BlogUseCase() IBlogUsecase {
	return c.connections[_blogUseCase].(IBlogUsecase)
}
