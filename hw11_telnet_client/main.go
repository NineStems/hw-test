package main

// Go provides a `flag` package supporting basic
// command-line flag parsing. We'll use this package to
// implement our example command-line program.
import (
	"flag"
	"log"
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

	fmt.Println("timeout:", timeout.String())
	fmt.Println("host:", host)
	fmt.Println("port:", port)
}
