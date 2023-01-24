package main

import (
	"fmt"
	"os"

	"github.com/brettmostert/fnple/build/internal/commands"
	"github.com/brettmostert/fnple/pkg/errors/exitError"
)

func main() {
	e := commands.NewExecuter()
	if _, err := e.Execute(); err != nil {
		fmt.Printf("failed executing command with error (%v)\n", err)
		os.Exit(int(exitError.Failure))
	}
}
