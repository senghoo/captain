package command

import (
	"bytes"
	"os"

	"github.com/fsouza/go-dockerclient"
	"github.com/senghoo/captain/models"
)

type DockerBuildZipCommand struct {
	DockerID int64
	File     string
	Name     string
	buffer   *bytes.Buffer
	next     Command
}

const defaultBuffSize = 1024 * 1024

func NewDockerBuildZipCommand(name, file string, dockerID int64) *DockerBuildZipCommand {
	buf := make([]byte, 0, defaultBuffSize)
	return &DockerBuildZipCommand{
		DockerID: dockerID,
		Name:     name,
		File:     file,
		buffer:   bytes.NewBuffer(buf),
	}
}

func (d *DockerBuildZipCommand) Run(build *models.Build) {
	logger := build.Logger()
	d, err := models.GetDockerServerByID(d.DockerID)
	if err != nil {
		logger.Printf("Error: %s", err)
		return
	}
	// file
	file, err := os.Open(d.File)
	if err != nil {
		logger.Printf("Error: %s", err)
		return
	}
	d.Build(docker.BuildImageOptions{
		Name:        d.Name,
		InputSteam:  file,
		OutpubSteam: d.buffer,
	})
}

func (r *DockerBuildZipCommand) Next() Command {
	return r.next
}

func (r *DockerBuildZipCommand) SetNext(c Command) {
	r.next = c
}
