package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/codegangsta/cli"
	"github.com/senghoo/captain/cmd"
	"github.com/senghoo/captain/models"
)

const APP_VER = "0.0.1"
const ENGINE_MAX_RETRY = 3

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// init orm
	for i := 0; i < ENGINE_MAX_RETRY; i++ {
		if err := models.NewEngine(); err != nil {
			fmt.Print(err)
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	// start app
	app := cli.NewApp()
	app.Name = "Captain"
	app.Usage = "A Docker Manage Service"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdWeb,
		cmd.UserAdd,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
