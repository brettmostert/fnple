package commands

import (
	config "github.com/brettmostert/fnple/build/internal/builder"
	"github.com/brettmostert/fnple/pkg/cli"
)

func (e *Executer) initTest() {
	cmd := &cli.Command{
		Name:             "test",
		Run:              ExecuteTest,
		ShortDescription: "provides functionality to test projects",
		// Args: []string{"project"},
	}
	cmd.Args().Set("project")
	cmd.Flags().String("f", "build.json", "")

	e.rootCommand.AddCommand(cmd)
}

func ExecuteTest(cmd *cli.Command, args []string) ([]interface{}, error) {
	builder := config.NewBuilder(cmd.Flags().GetString("f"))
	err := builder.Test(cmd.Args().Get("project"))

	return nil, err
}
