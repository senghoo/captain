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
		stringFlag("listen, l", "localhost:4000", "Temporary port number to prevent conflict"),
	},
}

func runWeb(ctx *cli.Context) {
	server := web.NewServer()
	server.Run(ctx.String("listen"))
}
