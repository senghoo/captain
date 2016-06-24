package command

import (
	"bytes"
	"os"
	"path"

	"github.com/fsouza/go-dockerclient"
	"github.com/senghoo/captain/models"
)

type DockerBuildArchiveCommand struct {
	DockerID int64
	File     string
	Name     string
	buffer   *bytes.Buffer
	next     Command
}

const defaultBuffSize = 1024 * 1024

func NewDockerBuildArchiveCommand(name, file string, dockerID int64) *DockerBuildArchiveCommand {
	buf := make([]byte, 0, defaultBuffSize)
	return &DockerBuildArchiveCommand{
		DockerID: dockerID,
		Name:     name,
		File:     file,
		buffer:   bytes.NewBuffer(buf),
	}
}

func (d *DockerBuildArchiveCommand) Run(build *models.Build) {
	logger := build.Logger()
	dockerServ, err := models.GetDockerServerByID(d.DockerID)
	if err != nil {
		logger.Printf("Error: %s", err)
		return
	}
	// file
	file, err := os.Open(path.Join(build.Path(), d.File))
	if err != nil {
		logger.Printf("Error: %s", err)
		return
	}
	err = dockerServ.Build(docker.BuildImageOptions{
		Name:         d.Name,
		InputStream:  file,
		OutputStream: d.buffer,
	})
	if err != nil {
		logger.Printf("Error: %s", err)
		return
	}

	logger.Printf("Output:\n>>>>>%s\n<<<<<", d.buffer.String())
}

func (d *DockerBuildArchiveCommand) Next() Command {
	return d.next
}

func (d *DockerBuildArchiveCommand) SetNext(c Command) {
	d.next = c
}
