package usecase

import (
	"database/sql"
	"news/internal/service/repo/postgres"
	"news/pkg/logger"
)

type IUseCase interface {
	NewsUseCase() INewsUsecase
}
type UseCase struct {
	connections map[string]interface{}
}

const (
	_newsUseCase = "news_use_case"
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
	return &UseCase{
		connections: connections,
	}
}
func (c *UseCase) NewsUseCase() INewsUsecase {
	return c.connections[_newsUseCase].(INewsUsecase)
}
