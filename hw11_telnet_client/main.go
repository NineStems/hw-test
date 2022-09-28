package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)
import "fmt"

var ErrNotEnoughArgs = fmt.Errorf("not enouth arguments for establisment connection")

func main() {
	timeout := flag.Duration("timeout", time.Second*10, "a duration")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		log.Fatal(ErrNotEnoughArgs)
	}

	host, port := args[0], args[1]
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	addr := net.JoinHostPort(host, port)

	client := NewTelnetClient(addr, *timeout, os.Stdin, os.Stdout)
	defer client.Close()

	if err := client.Connect(); err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	go func() {
		defer cancel()

		if err := client.Send(); err != nil {
			log.Fatal(fmt.Errorf("client.Send: %w", err))
			return
		}
	}()

	go func() {
		defer cancel()

		if err := client.Receive(); err != nil {
			log.Fatal(fmt.Errorf("client.Receive: %w", err))
			return
		}
	}()

	<-ctx.Done()
}
