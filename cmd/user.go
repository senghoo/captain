package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/senghoo/captain/models"
)

var UserAdd = cli.Command{
	Name:        "useradd",
	Usage:       "useradd  username password [-admin]",
	Description: `Add user.`,
	Action:      addUser,
	Flags: []cli.Flag{
		boolFlag("admin", "set as adminstrator"),
	},
}

func addUser(ctx *cli.Context) {
	username := ctx.Args()[0]
	password := ctx.Args()[1]

	if username == "" || password == "" {
		fmt.Println("username and password required")
	}

	u := models.NewUser(username)
	u.SetPassword(password)
	u.Save()
}
