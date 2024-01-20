package postgres

import (
	"database/sql"
	"fmt"
	
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"
	configs "news/internal/config"
	"sync"
)

var (
	instance *sql.DB
	once sync.Once
)
func DB(cfg *configs.Postgres) (*sql.DB,error){
	var err error
	once.Do(func(){
		psqlString := fmt.Sprintf(
			`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.Password,
			cfg.Database,
			)
		instance, err = sql.Open("pgx", psqlString)
	})
	
	if err != nil {
		return nil, errors.Wrap(err, "pgx.Connect")
	}
	return instance,nil
}