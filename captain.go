package main

import (
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/gogits/gogs/models"
	"github.com/senghoo/captain/cmd"
)

const APP_VER = "0.0.1"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// init orm
	models.NewEngine()

	// start app
	app := cli.NewApp()
	app.Name = "Captain"
	app.Usage = "A Docker Manage Service"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdWeb,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
