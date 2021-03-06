package command

import (
	"bytes"
	"os"
	"path"

	"github.com/fsouza/go-dockerclient"
	"github.com/senghoo/captain/models"
)

func init() {
	RegisterCommand("DockerBuildArchiveCommand", new(DockerBuildArchiveCommand))
}

type DockerBuildArchiveCommand struct {
	DockerID int64
	File     string
	Name     string
}

const defaultBuffSize = 1024 * 1024

func NewDockerBuildArchiveCommand(name, file string, dockerID int64) *DockerBuildArchiveCommand {
	return &DockerBuildArchiveCommand{
		DockerID: dockerID,
		Name:     name,
		File:     file,
	}
}

func (d *DockerBuildArchiveCommand) Run(build *models.Build) string {
	logger := build.Logger()
	buffer := bytes.NewBuffer(nil)
	dockerServ, err := models.GetDockerServerByID(d.DockerID)
	if err != nil {
		logger.Printf("Error: %s", err)
		return "error"
	}
	// file
	file, err := os.Open(path.Join(build.Path(), d.File))
	if err != nil {
		logger.Printf("Error: %s", err)
		return "error"
	}
	err = dockerServ.Build(docker.BuildImageOptions{
		Name:         d.Name,
		InputStream:  file,
		OutputStream: buffer,
	})
	if err != nil {
		logger.Printf("Error: %s", err)
		return "error"
	}

	logger.Printf("Output:\n>>>>>%s\n<<<<<", buffer.String())
	return "success"
}
