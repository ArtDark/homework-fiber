package database

import (
	"context"
	"homework-fiber/config"
	"homework-fiber/pkg/logger"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDbPool(config *config.DatabaseConfig, logger *logger.Service) *pgxpool.Pool {	
	pool, err := pgxpool.New(context.Background(), config.DbUrl())

	if err != nil {
		logger.Error.Error("не удалось подключиться к БД")
		os.Exit(1)
	}
	logger.Info.Info("подключение к БД установлено")
	return pool
}

func CloseDbPool(pool *pgxpool.Pool, logger *logger.Service) {
	pool.Close()
	logger.Info.Info("подключение к БД закрыто")
}