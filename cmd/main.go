package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/emp1re/students/cmd/handler"
	"github.com/emp1re/students/internal/config"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", cfg.Database.User, cfg.Database.Pass, cfg.Database.Name)
	//connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s host=172.17.0.2 port=5432 sslmode=disable", cfg.Database.User, cfg.Database.Pass, cfg.Database.Name)
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	if err := conn.Ping(context.Background()); err != nil {

		return fmt.Errorf("Unable to ping database:%s", zap.Error(err))
	}
	logger.Info("Connected to POSTGRESQL")

	handlers := handler.NewHandler(nil, *logger)
	router := handler.MakeRouter(handlers)
	err = router.Run(":8080")
	if err != nil {
		return fmt.Errorf("router.Run: %w", zap.Error(err))
	}
	return nil
}
