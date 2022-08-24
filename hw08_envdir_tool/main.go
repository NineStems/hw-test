package main

import (
	"fmt"
	"log"
	"os"
)

var ToLittleArgs = fmt.Errorf("count of cmd arguments is less than 2")

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatal(ToLittleArgs)
	}
	envs, err := ReadDir(args[0])
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(RunCmd(args[1:], envs))
}
