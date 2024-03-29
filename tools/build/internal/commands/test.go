package commands

import (
	"fmt"

	"github.com/brettmostert/fnple-go/go/pkg/cli"
	config "github.com/brettmostert/fnple-go/tools/build/internal/builder"
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
	fmt.Printf("%s\n", args)
	err := builder.Test(cmd.Args().Get("project"))

	return nil, err
}
