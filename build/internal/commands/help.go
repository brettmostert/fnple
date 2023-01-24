package commands

import (
	"github.com/brettmostert/fnple/pkg/cli"
)

func (e *Executer) initHelp() {
	cmd := &cli.Command{
		Name: "help",
		Run:  ExecuteHelp,
	}
	cmd.Args().Set("command", "the command for which you need help")
	e.rootCommand.AddCommand(cmd)
}

func ExecuteHelp(cmd *cli.Command, args []string) ([]interface{}, error) {
	cmd.PrintHelpText()

	return nil, nil
}
