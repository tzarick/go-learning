package main

import (
	"flag"

	"github.com/tzarick/go-learning/big-picture/simple-cli/cli"
)

func main() {
	path := flag.String("path", "testLog.log", "The path to the log that should be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for")
	flag.Parse()

	cli.Run_cli(path, level)
}
