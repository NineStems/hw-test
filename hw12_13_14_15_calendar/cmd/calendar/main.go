package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq" //nolint: gci

	"github.com/hw-test/hw12_13_14_15_calendar/internal/config"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/database/postgres"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/interactor/app"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/pkg/logger"
	memorystorage "github.com/hw-test/hw12_13_14_15_calendar/internal/repository/storage/memory"
	sqlstorage "github.com/hw-test/hw12_13_14_15_calendar/internal/repository/storage/sql"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/server"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/calendar_config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	cfg := config.New()
	err := cfg.Apply(configFile)
	if err != nil {
		panic(err)
	}

	log := logger.Console(cfg.Logger.Path, cfg.Logger.Level)
	sugarLog := logger.InitSugarZapLogger(log)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	// INFO: не самый лучший подход, но из-за разницы в хранилищах сходу не придумал, как это обернуть
	// Идея использовать методы Open/Close у репы решит проблему, но управлять конектом на этом слое - тоже не выход
	var storage app.Storage
	switch cfg.Database.Source {
	case "postgres":
		db := postgres.New(&cfg.Database, sugarLog)
		err = db.Open(ctx)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		storage = sqlstorage.New(db)
	case "inmemory":
		storage = memorystorage.New(sugarLog)
	default:
		panic(fmt.Sprintf("undefined database source=%v", cfg.Database.Source))
	}

	calendar := app.New(sugarLog, storage)

	srv := server.NewServer(sugarLog, cfg, calendar)
	if err = srv.Start(ctx); err != nil {
		sugarLog.Error(err)
	}
}
