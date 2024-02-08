package main

import (
	"context"
	"fmt"
	"time"

	"github.com/emp1re/students/cmd/client"
	"github.com/emp1re/students/cmd/handler"
	"github.com/emp1re/students/internal/config"
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
	conn, err := client.ConnectDb(ctx, *cfg, logger)
	if err != nil {
		return fmt.Errorf("ConnectDb: %w", zap.Error(err))
	}
	fmt.Println(conn)

	handlers := handler.NewHandler(nil, *logger)
	router := handler.MakeRouter(handlers)
	err = router.Run(":8030")
	if err != nil {
		return fmt.Errorf("router.Run: %w", zap.Error(err))
	}
	return nil
}
