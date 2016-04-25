package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/senghoo/captain/web"
)

var CmdWeb = cli.Command{
	Name:        "web",
	Usage:       "Start Captain web server",
	Description: `Captain web server.`,
	Action:      runWeb,
	Flags: []cli.Flag{
		stringFlag("port, p", "8000", "Temporary port number to prevent conflict"),
	},
}

func runWeb(ctx *cli.Context) {
	server := web.NewServer()
	server.Run()
}
