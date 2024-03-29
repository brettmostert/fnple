package commands

import (
	"github.com/brettmostert/fnple-go/go/pkg/cli"
	config "github.com/brettmostert/fnple-go/tools/build/internal/builder"
)

func (e *Executer) initBuild() {
	cmd := &cli.Command{
		Name:             "build",
		Run:              ExecuteBuild,
		ShortDescription: "provides build functionality",
		// Args: []string{"project"},
	}
	cmd.Args().Set("project")
	cmd.Flags().String("f", "build.json", "")
	cmd.Flags().String("release", "", "")

	e.rootCommand.AddCommand(cmd)
}

func ExecuteBuild(cmd *cli.Command, args []string) ([]interface{}, error) {
	builder := config.NewBuilder(cmd.Flags().GetString("f"))
	err := builder.Build(cmd.Args().Get("project"), cmd.Flags().GetString("release"))

	return nil, err
}
