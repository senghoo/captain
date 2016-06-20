package command

import "github.com/senghoo/captain/models"

type DockerBuildZipCommand struct {
	DockerID int64
	File     string
	next     Command
}

func NewDockerBuildZipCommand(file string, dockerID int64) *DockerBuildZipCommand {
	return &DockerBuildZipCommand{
		DockerID: dockerID,
		File:     file,
	}
}

func (d *DockerBuildZipCommand) Run(build *models.Build) {
	docker, err := models.GetDockerServerByID(d.DockerID)
	if err != nil {
		logger.Printf("Error: %s", err)
		return
	}

}

func (r *DockerBuildZipCommand) Next() Command {
	return r.next
}

func (r *DockerBuildZipCommand) SetNext(c Command) {
	r.next = c
}
