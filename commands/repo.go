package command

import (
	"errors"
	"path"

	"github.com/senghoo/captain/models"
)

type RepoUpdateCommand struct {
	RepoID int64
	Status int
	next   Command
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

func (r *RepoUpdateCommand) SetArgs(args CommandArgs) error {
	repoID, ok := args.Int64("RepoID")
	if !ok {
		return errors.New("RepoID not set")
	}
	r.RepoID = repoID
	return nil
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
	Status  int
	Format  string
	Branch  string
	OutFile string
	next    Command
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

func (r *RepoArchiveCommand) SetArgs(args CommandArgs) error {
	repoID, ok := args.Int64("RepoID")
	if !ok {
		return errors.New("RepoID not set")
	}
	r.RepoID = repoID

	format, ok := args.String("Format")
	if !ok {
		return errors.New("Format not set")
	}
	r.Format = format

	branch, ok := args.String("Branch")
	if !ok {
		return errors.New("Branch not set")
	}
	r.Branch = branch

	file, ok := args.String("OugFile")
	if !ok {
		return errors.New("File not set")
	}
	r.OutFile = file

	return nil
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
}
