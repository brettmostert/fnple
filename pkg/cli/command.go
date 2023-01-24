package cli

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/brettmostert/fnple/pkg/errors/exitError"
)

type Command struct {
	Name             string
	ShortDescription string
	LongDescription  string
	Example          string
	Run              func(cmd *Command, args []string) ([]interface{}, error)
	HelpType         Help
	argSet           *ArgSet
	flagSet          *FlagSet
	// flags   []*string

	parent   *Command
	commands []*Command
}

type Options struct {
	Args []string
}

func (cmd *Command) PrintHelpText() {
	description := cmd.ShortDescription
	if len(description) > 0 {
		description = " - " + description
	}

	fmt.Printf("%s %s\n\n", cmd.Name, description)

	usageText := ""
	parent := cmd.parent

	for parent != nil {
		usageText = " " + parent.Name
		parent = parent.parent
	}

	usageText = usageText + " " + cmd.Name
	if len(cmd.commands) > 0 {
		usageText = usageText + " <command>"
	}

	fmt.Println("Usage:")
	fmt.Printf("%-5s %s\n\n", "", strings.TrimSpace(usageText)+" [args]")

	if len(cmd.argSet.args) > 0 {
		fmt.Printf("The args are:\n")
	}

	for _, arg := range cmd.argSet.args {
		fmt.Printf("%-5s %-10s %-5s\n", "", arg.Name, arg.description)
	}

	if len(cmd.commands) > 0 {
		fmt.Printf("The commands are:\n")
	}

	for _, subCmd := range cmd.commands {
		fmt.Printf("%-5s %-10s %-5s\n", "", subCmd.Name, subCmd.ShortDescription)
	}
}

func (cmd *Command) Execute(options ...Options) ([]interface{}, error) {
	args := os.Args[1:]

	if len(options) > 0 {
		args = options[0].Args
	}

	argsLen := len(args)

	var cmdToExecute *Command

	var argsToExecute []string

	switch {
	case argsLen == 0:
		cmdToExecute = cmd
		argsToExecute = args
	case argsLen > 0:
		cmdToExecute = cmd.findCommand(args)
		argsToExecute = args[argsLen-1:]
	}

	if cmdToExecute == nil {
		return nil, exitError.New("Command not found, args: "+strings.Join(args, " "), exitError.NotFound)
	}

	err := cmdToExecute.parseFlags(argsToExecute)
	if err != nil {
		return nil, exitError.New("Unable to parse flags, args: "+strings.Join(args, " "), exitError.InvalidFlags)
	}

	i := 0

	sort.Ints(cmdToExecute.Args().keys)

	for _, key := range cmdToExecute.Args().keys {
		if i > len(argsToExecute)-1 {
			break
		}

		cmdToExecute.Args().args[key].value = argsToExecute[i]

		i++
	}

	return cmdToExecute.Run(cmdToExecute, argsToExecute)
}

func (cmd *Command) findNext(args []string) *Command {
	var nextCommand *Command

	for _, cmd := range cmd.commands {
		if cmd.Name == args[0] {
			nextCommand = cmd
			break
		}
	}

	return nextCommand
}

func (rootCmd *Command) findCommand(args []string) *Command {
	var command *Command

	var innerfind func(*Command, []string) *Command

	innerfind = func(innerCommand *Command, innerArgs []string) *Command {
		if len(innerArgs) == 0 {
			return innerCommand
		}

		command = innerCommand.findNext(innerArgs)
		if command == nil {
			return innerCommand
		}

		if command != nil {
			command = innerfind(command, innerArgs[1:])
		}

		return command
	}

	parentCmd := rootCmd.findNext(args)
	if parentCmd != nil {
		argsWithoutFlags := stripFlags(args[1:])
		command = innerfind(parentCmd, argsWithoutFlags)

		if command == nil {
			if len(parentCmd.Args().args) > 0 {
				command = parentCmd
			}
		}
	}

	return command
}

func (parentCmd *Command) AddCommand(cmds ...*Command) {
	for _, cmd := range cmds {
		if cmd == parentCmd {
			panic("Command cannot be a child of itself")
		}

		if parentCmd.Name == cmd.Name {
			panic("Parent command and child command cannot have the same name")
		}

		cmd.parent = parentCmd
		parentCmd.commands = append(parentCmd.commands, cmd)
	}
}
