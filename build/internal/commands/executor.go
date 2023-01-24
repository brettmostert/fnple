package commands

import (
	"github.com/brettmostert/fnple/pkg/cli"
)

type Executer struct {
	rootCommand *cli.Command
}

func NewExecuter() *Executer {
	rootCommand := &cli.Command{
		Name:             "bob",
		ShortDescription: "Bob the Builder",
		LongDescription:  "A builder builds what a builder builds",
		Example:          "bob build {projectName}",
		Run:              DefaultCmd,
		HelpType:         cli.Default,
	}

	e := &Executer{
		rootCommand: rootCommand,
	}

	e.initConfig()
	e.initBuild()
	e.initTest()
	e.initProject()
	e.initTern()

	//TODO: Implement help generically
	// e.initHelp()

	return e
}

func (e *Executer) Execute() (interface{}, error) {
	return e.rootCommand.Execute()
}

func DefaultCmd(cmd *cli.Command, args []string) ([]interface{}, error) {
	cmd.PrintHelpText()
	return nil, nil
}
