package commands

import (
	config "github.com/brettmostert/fnple/build/internal/builder"
	"github.com/brettmostert/fnple/pkg/cli"
)

func (e *Executer) initConfig() {
	cmd := &cli.Command{
		Name:             "config",
		Run:              ExecuteConfig,
		ShortDescription: "provides functionality for the config",
	}

	subCmd := &cli.Command{
		Name:             "print",
		ShortDescription: "prints out the configuration file",
		Run:              ExecutePrint,
	}

	subCmd.Flags().String("f", "build.json", "")

	cmd.AddCommand(subCmd)
	e.rootCommand.AddCommand(cmd)
}

func ExecuteConfig(cmd *cli.Command, args []string) ([]interface{}, error) {
	cmd.PrintHelpText()
	return nil, nil
}

func ExecutePrint(cmd *cli.Command, args []string) ([]interface{}, error) {
	config.Print(cmd.Flags().GetString("f"))
	return nil, nil
}
