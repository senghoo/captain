package command

import (
	"path"

	"github.com/senghoo/captain/models"
)

func init() {
	RegisterCommand("RepoUpdateCommand", new(RepoUpdateCommand))
	RegisterCommand("RepoArchiveCommand", new(RepoArchiveCommand))
}

type RepoUpdateCommand struct {
	RepoID int64
}

const (
	Success = iota
	WorkspaceNotFound
	RepoNotFound
)

func (r *RepoUpdateCommand) Clone() Command {
	n := *r
	return &n
}

func (r *RepoUpdateCommand) Run(build *models.Build) string {
	logger := build.Logger()

	repo := new(models.Repository)
	has, err := models.GetByID(r.RepoID, repo)
	if !has {
		logger.Printf("Error: %d has not exists", r.RepoID)
		return "error"
	}
	if err != nil {
		logger.Printf("Error: %s", err)
		return "error"
	}
	out, err := repo.Update()
	logger.Printf("Output:\n>>>>>%s\n<<<<<", out)
	if err != nil {
		logger.Printf("Error: %s", err)
		return "error"
	}
	return "success"
}

type RepoArchiveCommand struct {
	RepoID  int64
	Format  string
	Branch  string
	OutFile string
}

func NewRepoArchiveCommand(repoID int64, format, branch, file string) *RepoArchiveCommand {
	return &RepoArchiveCommand{
		RepoID:  repoID,
		Format:  format,
		Branch:  branch,
		OutFile: file,
	}
}

func (r *RepoArchiveCommand) Clone() Command {
	n := *r
	return &n
}

func (r *RepoArchiveCommand) Run(build *models.Build) string {
	logger := build.Logger()

	repo := new(models.Repository)
	has, err := models.GetByID(r.RepoID, repo)
	if !has {
		logger.Printf("Error: %d has not exists", r.RepoID)
		return "error"
	}
	if err != nil {
		logger.Printf("Error: %s", err)
		return "error"
	}
	p := path.Join(build.Path(), r.OutFile)
	out, err := repo.Archive(r.Format, r.Branch, p)
	logger.Printf("Output:\n>>>>>%s\n<<<<<", out)
	if err != nil {
		logger.Printf("Error: %s", err)
		return "error"
	}
	return "success"
}
