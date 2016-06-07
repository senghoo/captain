package command

import "github.com/senghoo/captain/models"

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

func NewRepoUpdateCommand(repoID int64) *RepoUpdateCommand {
	return &RepoUpdateCommand{
		RepoID: repoID,
	}
}

func (r *RepoUpdateCommand) Run(build *models.Build) {
	logger := build.Logger()

	repo := new(models.Repository)
	has, err := models.GetByID(r.RepoID, repo)
	if !has {
		logger.Printf("Error: %d has not exists", r.RepoID)
		return
	}
	if err != nil {
		logger.Printf("Error: %s", err)
		return
	}
	out, err := repo.Update()
	logger.Printf("Output:\n>>>>>%s\n<<<<<", out)
	if err != nil {
		logger.Printf("Error: %s", err)
		return
	}
}

func (r *RepoUpdateCommand) Next() Command {
	return r.next
}

func (r *RepoUpdateCommand) SetNext(c Command) {
	r.next = c
}

type RepoArchiveCommand struct {
	RepoID int64
	Status int
	next   Command
}

func NewRepoArchiveCommand(repoID int64) *RepoUpdateCommand {
	return &RepoUpdateCommand{
		RepoID: repoID,
	}
}

func (r *RepoArchiveCommand) Run(build *models.Build) {
	logger := build.Logger()

	repo := new(models.Repository)
	has, err := models.GetByID(r.RepoID, repo)
	if !has {
		logger.Printf("Error: %d has not exists", r.RepoID)
		return
	}
	if err != nil {
		logger.Printf("Error: %s", err)
		return
	}
	out, err := repo.Archive()
	logger.Printf("Output:\n>>>>>%s\n<<<<<", out)
	if err != nil {
		logger.Printf("Error: %s", err)
		return
	}
}

func (r *RepoArchiveCommand) Next() Command {
	return r.next
}

func (r *RepoArchiveCommand) SetNext(c Command) {
	r.next = c
}
