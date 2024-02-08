package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/emp1re/students/internal/config"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	ctx := context.Background()
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	logger, err := zap.NewProduction()

	url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", cfg.Database.User, cfg.Database.Pass, cfg.Database.Name)
	//connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s host=172.17.0.2 port=5432 sslmode=disable", cfg.Database.User, cfg.Database.Pass, cfg.Database.Name)
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	//if err := conn.Ping(context.Background()); err != nil {
	//
	//	return nil, fmt.Errorf("Unable to ping database:%s", zap.Error(err))
	//}
	logger.Info("Connected to POSTGRESQL")
	testQueries = New(conn)
	os.Exit(m.Run())
}
