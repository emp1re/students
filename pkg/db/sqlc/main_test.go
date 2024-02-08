package db

import (
	"context"
	"os"
	"testing"

	"github.com/emp1re/students/cmd/client"
	"github.com/emp1re/students/internal/config"
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

	conn, err := client.ConnectDb(ctx, *cfg, logger)
	if err != nil {
		panic(err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
