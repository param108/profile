package cmd

import "github.com/urfave/cli/v2"


var cmds = []*cli.Command{}

func GetCommands() []*cli.Command {
	return cmds
}

func register(cmd *cli.Command) {
	cmds = append(cmds, cmd)
}
