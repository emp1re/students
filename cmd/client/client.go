package client

import (
	"context"
	"fmt"
	"os"

	"github.com/emp1re/students/internal/config"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func ConnectDb(ctx context.Context, cfg config.Config, logger *zap.Logger) {
	url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", cfg.Database.User, cfg.Database.Pass, cfg.Database.Name)
	//connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s host=172.17.0.2 port=5432 sslmode=disable", cfg.Database.User, cfg.Database.Pass, cfg.Database.Name)
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		logger.Error("Unable to connect to database: %v\n", zap.Error(err))
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	if err := conn.Ping(context.Background()); err != nil {
		logger.Error("Unable to ping database", zap.Error(err))
		os.Exit(1)
	}
	logger.Info("Connected to POSTGRESQL")

}
