package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/senghoo/captain/cmd"
	"github.com/senghoo/captain/models"
)

const APP_VER = "0.0.1"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// init orm
	if err := models.NewEngine(); err != nil {
		fmt.Print(err)
		return
	}

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
