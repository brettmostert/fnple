package commands

import (
	"strconv"
	"strings"

	config "github.com/brettmostert/fnple/build/internal/builder"
	"github.com/brettmostert/fnple/pkg/cli"
)

func (e *Executer) initTern() {
	cmd := &cli.Command{
		Name:             "db",
		Run:              ExecuteTern,
		ShortDescription: "provides functionality to manage db's",
	}

	cmd.Args().Set("project")

	cmd.Args().Set("1")
	cmd.Args().Set("2")
	cmd.Args().Set("3")
	cmd.Args().Set("4")
	cmd.Args().Set("5")

	cmd.Flags().String("f", "build.json", "")

	e.rootCommand.AddCommand(cmd)
}

func ExecuteTern(cmd *cli.Command, args []string) ([]interface{}, error) {
	builder := config.NewBuilder(cmd.Flags().GetString("f"))

	var sb strings.Builder

	spacer := ""

	for i := 1; i <= cmd.Args().Length()-1; i++ {
		arg := cmd.Args().Get(strconv.Itoa(i))
		if arg != "" {
			sb.WriteString(spacer + arg)
			spacer = " "
		}
	}

	err := builder.Db(cmd.Args().Get("project"), sb.String())

	return nil, err
}
