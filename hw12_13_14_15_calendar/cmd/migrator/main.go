package main

import (
	"database/sql"
	"flag"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/pressly/goose/v3"

	"github.com/hw-test/hw12_13_14_15_calendar/internal/config"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/calendar_config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	cfg := config.New()
	err := cfg.Apply(configFile)
	if err != nil {
		panic(err)
	}

	source := fmt.Sprintf(
		"dbname=%v user=%v password=%v host=%v port=%v sslmode=disable",
		cfg.Database.Database,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port)

	var db *sql.DB
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", source)
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}
		break
	}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}
