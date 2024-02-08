package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/emp1re/students/cmd/handler"
	"github.com/emp1re/students/internal/config"
	"github.com/emp1re/students/internal/service"
	db "github.com/emp1re/students/pkg/db/sqlc"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		fmt.Errorf("run: %w", zap.Error(err))
		return
	}
}
func run() error {

	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()
	cfg, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("config.ReadConfig: %w", zap.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
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

	srvc := service.NewService(ctx, logger, conn, db.New(conn))
	handlers := handler.NewHandler(srvc, *logger)
	router := handler.MakeRouter(handlers)
	err = router.Run(":8030")
	if err != nil {
		return fmt.Errorf("router.Run: %w", zap.Error(err))
	}
	return nil
}
