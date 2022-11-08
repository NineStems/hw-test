package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os/signal"
	"syscall"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	v1 "github.com/hw-test/hw12_13_14_15_calendar/api/v1"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/config"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/interactor/sender"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/pkg/logger"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/repository/rabbit"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/sender_config.yaml", "Path to configuration file")
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

	var connRabbit *amqp.Connection
	retryTime := time.Second * 0
	retryCount := 0
	retryCountMax := 10

	for retryCount < retryCountMax && connRabbit == nil {
		time.Sleep(retryTime)
		retryTime += (retryTime + time.Second) * 2
		connRabbit, err = amqp.Dial(fmt.Sprintf(
			"amqp://%v:%v@%v/",
			cfg.Rabbit.Credential.Username,
			cfg.Rabbit.Credential.Password,
			net.JoinHostPort(cfg.Rabbit.Host, cfg.Rabbit.Port),
		))
		if err != nil {
			sugarLog.Error(err)
		}
		retryCount++
	}

	if connRabbit == nil {
		sugarLog.Error("Cant connect to rabbitMQ")
		return
	}

	defer connRabbit.Close()

	rabbitClient := rabbit.New(connRabbit, sugarLog, cfg.Rabbit)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	connGrpc, err := grpc.Dial(net.JoinHostPort(cfg.Server.Grpc.Host, cfg.Server.Grpc.Port), opts...)
	if err != nil {
		panic(err)
	}
	defer connGrpc.Close()

	client := v1.NewCalendarClient(connGrpc)

	app := sender.New(sugarLog, client, rabbitClient, cfg.Sender)
	err = app.Start(ctx)
	if err != nil {
		sugarLog.Error(err)
		return
	}
}
