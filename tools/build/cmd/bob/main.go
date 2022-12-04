package main

import (
	"fmt"
	"os"

	"github.com/brettmostert/fnple-go/go/pkg/errors/exitError"
	"github.com/brettmostert/fnple-go/tools/build/internal/commands"
)

func main() {
	e := commands.NewExecuter()
	if _, err := e.Execute(); err != nil {
		fmt.Printf("failed executing command with error (%v)\n", err)
		os.Exit(int(exitError.Failure))
	}
}
