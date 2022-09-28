package main

// Go provides a `flag` package supporting basic
// command-line flag parsing. We'll use this package to
// implement our example command-line program.
import (
	"context"
	"flag"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"
)
import "fmt"

var NotEnoughArgs = fmt.Errorf("not enouth arguments for establisment connection")

func main() {
	timeout := flag.Duration("timeout", time.Second*10, "a duration")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		log.Fatal(NotEnoughArgs)
	}

	host, port := args[0], args[1]
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	addr := net.JoinHostPort(host, port)

	client := NewTelnetClient(addr, *timeout, nil, nil)
	defer client.Close()

	err := client.Connect()
	if err != nil {
		log.Fatal(fmt.Errorf("client.Connect: %w", err))
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
