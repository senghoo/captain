package command

import "github.com/senghoo/captain/models"

type Command interface {
	Run(build *models.Build)
	Next() Command
}

func RunCommand(cmd Command, build *models.Build) {
	for ; cmd != nil; cmd = cmd.Next() {
		cmd.Run(build)
	}
}
