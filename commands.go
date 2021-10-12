package main

import (
	"fmt"
	"os"
	"togo/command"

	"github.com/urfave/cli"
)

//global option for all command.
var GlobalFlags = []cli.Flag{}

var forceFlags = []cli.Flag{
	cli.BoolFlag{
		Name: "force, f",
		Usage: "Force initialize even if database already existed.",
	},
}

var allListFlags = []cli.Flag{
	cli.BoolFlag{
		Name: "all, a",
		Usage: "List all tasks.",
	},
}

var allDeleteFlagas = []cli.Flag{
	cli.BoolFlag{
		Name: "all, a",
		Usage: "Delete all tasks.",
	},
}

//commands are slice of command.
var Commands = []cli.Command{
	{
		Name: "init",
		Usage: "Initialize SQlite database",
		Action: command.CmdInit,
		Flags: forceFlags,
	},
	{
		Name: "add",
		Aliases: []string{"a"},
		Usage: "Add task",
		Action: command.CmdAdd,
		Flags: []cli.Flag{},
	},
	{
		Name: "update",
		Aliases: []string{"u"},
		Usage: "Update task",
		Action: command.CmdUpdate,
		Flags: []cli.Flag{},
	},
	{
		Name: "done",
		Aliases: []string{"d"},
		Usage: "Mark a task as done",
		Action: command.CmdDone,
		Flags: []cli.Flag{},
	},
	{
		Name: "list",
		Aliases: []string{"l"},
		Usage: "Show all tasks",
		Action: command.CmdList,
		Flags: allListFlags,
	},
	{
		Name: "delete",
		Usage: "Delete task",
		Action: command.CmdDelete,
		Flags: allDeleteFlagas,
	},
}

//commandNotFound is custom error.
func CommandNotFound(ctx *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", ctx.App.Name, command, ctx.App.Name, ctx.App.Name)
	os.Exit(2)
}
